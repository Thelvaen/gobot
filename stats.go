package main

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12"
	"github.com/thelvaen/gobot/config"
	"github.com/thelvaen/gobot/models"
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

	var stats models.Stat

	stats.User = message.User.Name
	err := dataStore.Where("user = ?", stats.User).First(&stats).Error
	if err == gorm.ErrRecordNotFound {
		stats.Score = 1
		dataStore.Create(&stats)
	}
	if err == nil {
		stats.Score++
		dataStore.Save(&stats)
	}

	return ""
}

func getCliStats(message twitch.PrivateMessage) string {
	// Outputting stats to the channel
	if message.Channel != config.Cred.Channel {
		return ""
	}

	var stats models.Stat

	stats.User = message.User.Name
	err := dataStore.Where("user = ?", stats.User).First(&stats).Error
	if err != nil {
		return ""
	}
	return "Ton score est : " + strconv.Itoa(stats.Score)
}

func getStats(ctx iris.Context) {
	var stats []models.Stat
	data := map[string]map[string]string{
		"Statistiques": {},
	}

	result := dataStore.Find(&stats)

	if result.Error == nil {
		for _, row := range stats {
			data["Statistiques"][row.User] = strconv.Itoa(row.Score)
		}
	}
	ctx.ViewData("Data", data)
	if err := ctx.View("stats.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
