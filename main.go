package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Thelvaen/gobot/config"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12"
)

func main() {
	// Initializing twitch Client
	if config.IsAuth {
		// Credential provided, authed connection
		twitchC = twitch.NewClient(config.Cred.Channel, config.Cred.Token)
	} else {
		// No credentials provided, anon connection
		twitchC = twitch.NewAnonymousClient()
	}
	// Registering Twitch IRC Client callback functions
	twitchC.OnPrivateMessage(func(message twitch.PrivateMessage) {
		parseMessage(message)
	})

	// Initializing modules
	initAggregator()
	initDice()
	initStats()

	// Initializing WebServer for the bot
	app := webBot()

	// Executing WebServer for the bot
	if config.WebConf.IsSecure {
		go app.Run(iris.TLS(config.WebConf.IP+":"+config.WebConf.Port, config.WebConf.Cert, config.WebConf.Key))
		p, _ := strconv.Atoi(config.WebConf.Port)
		newPort := strconv.Itoa(p + 1)
		srv1 := &http.Server{Addr: config.WebConf.IP + ":" + newPort, Handler: app}
		go srv1.ListenAndServe()
	} else {
		go app.Listen(config.WebConf.IP + ":" + config.WebConf.Port)
	}

	// Telling TwitchBot to join Main Channel
	twitchC.Join(config.Cred.Channel)

	// Connection to Twitch
	err := twitchC.Connect()
	if err != nil {
		log.Fatalf("can't connect to Twitch: %s", err)
	}
}
