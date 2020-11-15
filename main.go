package main

import (
	"regexp"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
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
		for filter, filterDetails := range Filters {
			found, _ := regexp.MatchString(filter, message.Message)
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

	// Initializing modules needs to be done after TwitchConnect
	initAggregator()
	initDice()
	initGiveAway()
	initPolls()

	// Starting web server as a Go Routine (background thread)
	server = gin.New()

	// Setting server in production mode
	//gin.SetMode(gin.ReleaseMode)

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	server.Use(gin.Recovery())

	// Passing templates to gin-gonic web server
	server.HTMLRender = createMyRender()

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
