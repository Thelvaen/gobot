package main

import (
	"regexp"

	"github.com/Thelvaen/gobot/models"
	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/gorm"
)

func pushAndSay(data string) {
	twitchC.Say(conf.Cred.Channel, data)
}

func parseMessage(message twitch.PrivateMessage) {
	if conf.IsAuth {
		// Command to process
		updateTwitchUser(message.User)
		for _, filterDetails := range filters {
			found, _ := regexp.MatchString(filterDetails.filterRegEx, message.Message)
			if found {
				botProcess := filterDetails.filterFunc(message)
				if message.Channel == conf.Cred.Channel {
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
	if err == gorm.ErrRecordNotFound {
		u.Name = user.Name
		u.DisplayName = user.DisplayName
		u.Statistique = models.Stat{
			Score: 0,
		}
		dataStore.Create(&u)
	}
	if u.Name != user.Name {
		u.Name = user.Name
		u.DisplayName = user.DisplayName
		dataStore.Save(&u)
	}
	if u.DisplayName != user.DisplayName {
		u.Name = user.Name
		u.DisplayName = user.DisplayName
		dataStore.Save(&u)
	}
}
