package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	statement *sql.Stmt
)

// Stats structure made exportable to be used with Gorm ORM
type Stats struct {
	gorm.Model
	User  string `gorm:"not null;unique"` // Utilisateur unique!
	Score int    `gorm:"not null;"`
}

func initStats() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  pushStats,
		FilterRegEx: ".*",
	})

	Filters = append(Filters, CLIFilter{
		FilterFunc:  getCliStats,
		FilterRegEx: "^!score$",
	})

	BotConfig.DataStore.AutoMigrate(&Stats{})
}

func pushStats(message twitch.PrivateMessage) string {
	if message.Channel != BotConfig.Cred.Channel {
		return ""
	}
	if len(message.Message) < 10 {
		return ""
	}
	if strings.HasPrefix(message.Message, "!") {
		return ""
	}

	var stats Stats

	stats.User = message.User.Name
	err = BotConfig.DataStore.Where("user = ?", stats.User).First(&stats).Error
	if err == gorm.ErrRecordNotFound {
		stats.Score = 1
		BotConfig.DataStore.Create(&stats)
	}
	if err == nil {
		stats.Score++
		BotConfig.DataStore.Save(&stats)
	}

	return ""
}

func getCliStats(message twitch.PrivateMessage) string {
	// Outputting stats to the channel
	if message.Channel != BotConfig.Cred.Channel {
		return ""
	}

	var stats Stats

	stats.User = message.User.Name
	err = BotConfig.DataStore.Where("user = ?", stats.User).First(&stats).Error

	return "Ton score est : " + strconv.Itoa(stats.Score)
}

func getStats(c *gin.Context) {
	var stats []Stats
	data := map[string]map[string]string{
		"Statistiques": {},
	}

	result := BotConfig.DataStore.Find(&stats)

	if result.Error == nil {
		for _, row := range stats {
			data["Statistiques"][row.User] = strconv.Itoa(row.Score)
		}
	}
	c.HTML(http.StatusOK, "stats.html", gin.H{
		"Context": prepareContext(c),
		"Data":    data,
	})
}
