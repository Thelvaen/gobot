package main

import (
	"log"
	"strconv"

	"github.com/Thelvaen/gobot/config"
	"github.com/gempir/go-twitch-irc/v2"
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
	go app.Listen(config.WebConf.URL + ":" + strconv.Itoa(config.WebConf.Port))

	// Telling TwitchBot to join Main Channel
	twitchC.Join(config.Cred.Channel)

	// Connection to Twitch
	err := twitchC.Connect()
	if err != nil {
		log.Fatalf("can't connect to Twitch: %s", err)
	}
}
