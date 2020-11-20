package main

import (
	"regexp"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/thelvaen/gobot/config"
)

func pushAndSay(data string) {
	twitchC.Say(config.Cred.Channel, data)
}

func parseMessage(message twitch.PrivateMessage) {
	if config.IsAuth {
		// Command to process
		for _, filterDetails := range filters {
			found, _ := regexp.MatchString(filterDetails.filterRegEx, message.Message)
			if found {
				botProcess := filterDetails.filterFunc(message)
				if message.Channel == config.Cred.Channel {
					pushAndSay(botProcess)
				}
			}
		}
	}
}
