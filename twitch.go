package main

import (
	"regexp"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/thelvaen/gobot/config"
	"github.com/thelvaen/gobot/models"
	"gorm.io/gorm"
)

func pushAndSay(data string) {
	twitchC.Say(config.Cred.Channel, data)
}

func parseMessage(message twitch.PrivateMessage) {
	if config.IsAuth {
		// Command to process
		updateTwitchUser(message.User)
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

func updateTwitchUser(user twitch.User) {
	u := models.TwitchUser{
		TwitchID: user.ID,
	}

	err := dataStore.Where(&u).First(&u).Error

	u.Name = user.Name
	u.DisplayName = user.DisplayName
	if err == gorm.ErrRecordNotFound {
		u.Statistique = models.Stat{
			Score: 0,
		}
		dataStore.Create(&u)
	} else {
		dataStore.Save(&u)
	}
}
