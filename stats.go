package main

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/models"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	statement *sql.Stmt
)

func initStats() {
	filters = append(filters, filter{
		filterFunc:  pushStats,
		filterRegEx: ".*",
	})

	filters = append(filters, filter{
		filterFunc:  getCliStats,
		filterRegEx: "^!score$",
	})
}

func pushStats(message twitch.PrivateMessage) string {
	if message.Channel != config.Cred.Channel {
		return ""
	}
	if len(message.Message) < 10 {
		return ""
	}
	if strings.HasPrefix(message.Message, "!") {
		return ""
	}

	var user models.TwitchUser

	err := dataStore.Preload("Statistique").Where("twitch_id = ?", message.User.ID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		updateTwitchUser(message.User)
	}
	if err == nil {
		user.Statistique.Score++
		dataStore.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)
	}

	return ""
}

func getCliStats(message twitch.PrivateMessage) string {
	// Outputting stats to the channel
	if message.Channel != config.Cred.Channel {
		return ""
	}

	var user models.TwitchUser

	err := dataStore.Preload("Statistique").Where("twitch_id = ?", message.User.ID).First(&user).Error
	if err != nil {
		return ""
	}
	return "Ton score est : " + strconv.Itoa(user.Statistique.Score)
}

func getStats(ctx iris.Context) {
	var users []models.TwitchUser
	data := map[string]map[string]string{
		"Statistiques": {},
	}

	result := dataStore.Preload("Statistique").Find(&users)

	if result.Error == nil {
		for _, row := range users {
			if row.Statistique.Score > 0 {
				data["Statistiques"][row.DisplayName] = strconv.Itoa(row.Statistique.Score)
			}
		}
	}
	ctx.ViewData("Data", data)
	if err := ctx.View("stats.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
