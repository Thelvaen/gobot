package main

import (
	"regexp"
	"strconv"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/thelvaen/gobot/csrf"
)

var (
	err error
)

func pushAndSay(data string) {
	BotConfig.TwitchC.Say(BotConfig.Cred.Channel, data)
}

func parseMessage(message twitch.PrivateMessage) {
	if BotConfig.Cred.IsAuth {
		// Command to process
		for _, filterDetails := range Filters {
			found, _ := regexp.MatchString(filterDetails.FilterRegEx, message.Message)
			if found {
				botProcess := filterDetails.FilterFunc(message)
				if message.Channel == BotConfig.Cred.Channel {
					pushAndSay(botProcess)
				}
			}
		}
	}
}

func main() {
	// Connecting to Twitch
	if BotConfig.Cred.IsAuth {
		BotConfig.TwitchC = twitch.NewClient(BotConfig.Cred.Channel, BotConfig.Cred.Token)
	} else {
		// No credentials provided, anon connection
		BotConfig.TwitchC = twitch.NewAnonymousClient()
	}
	// Registering Twitch IRC Client callback functions
	BotConfig.TwitchC.OnPrivateMessage(func(message twitch.PrivateMessage) {
		parseMessage(message)
	})

	// Setting server in production mode
	gin.SetMode(gin.ReleaseMode)

	// Starting web server as a Go Routine (background thread)
	server = gin.New()

	//store := cookie.NewStore([]byte(BotConfig.Cred.Channel + strconv.Itoa(time.Now().Nanosecond())))
	store := cookie.NewStore([]byte(BotConfig.Cred.Channel))
	server.Use(sessions.Sessions("mysession", store))

	server.Use(csrf.Middleware(csrf.Options{
		Secret: BotConfig.Cred.Channel + strconv.Itoa(time.Now().Nanosecond()),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	server.Use(gin.Recovery())

	// Passing templates to gin-gonic web server
	server.HTMLRender = createMyRender()

	// Initializing modules needs to be done after TwitchConnect
	initAuth()
	initAggregator()
	initDice()
	initGiveAway()
	initPolls()
	initStats()

	server.Use(static.Serve("/static", binaryFS("")))
	//http.Handle("/static", http.FileServer(AssetFile()))

	// Parsing routes to the server
	initRoutes()

	url := BotConfig.BotServer.URL + ":" + strconv.Itoa(BotConfig.BotServer.Port)
	go server.Run(url)

	// Telling TwitchBot to join Main Channel
	BotConfig.TwitchC.Join(BotConfig.Cred.Channel)

	// Connection to Twitch
	BotConfig.TwitchC.Connect()
	if err != nil {
		myPanic("can't connect to Twitch: %s", err)
	}
}
