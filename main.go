package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12"
)

func main() {
	// Initializing twitch Client
	if conf.IsAuth {
		// Credential provided, authed connection
		twitchC = twitch.NewClient(conf.Cred.Channel, conf.Cred.Token)
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
	if conf.WebConf.IsSecure {
		go app.Run(iris.TLS(conf.WebConf.IP+":"+conf.WebConf.Port, conf.WebConf.Cert, conf.WebConf.Key))
		p, _ := strconv.Atoi(conf.WebConf.Port)
		newPort := strconv.Itoa(p + 1)
		srv1 := &http.Server{Addr: conf.WebConf.IP + ":" + newPort, Handler: app}
		go srv1.ListenAndServe()
	} else {
		go app.Listen(conf.WebConf.IP + ":" + conf.WebConf.Port)
	}

	// Telling TwitchBot to join Main Channel
	twitchC.Join(conf.Cred.Channel)

	// Connection to Twitch
	err := twitchC.Connect()
	if err != nil {
		log.Fatalf("can't connect to Twitch: %s", err)
	}
}
