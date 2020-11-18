package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
// Block var local
)

// GiveAway structure made exportable to be used with Gorm ORM
type GiveAway struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string
	Status      bool
	Users       []User
}

func initGiveAway() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  registerGiveAway,
		FilterRegEx: "^!jeveux$",
	})

	BotConfig.DataStore.AutoMigrate(&GiveAway{})
}

func registerGiveAway(message twitch.PrivateMessage) (outMessage string) {
	outMessage = "toto"
	return
}

func getGiveAwayListForm(c *gin.Context) {
	c.String(200, "test")
}

func postGiveAwayList(c *gin.Context) {
	c.String(200, "test")
}

func getGiveAwayForm(c *gin.Context) {
	c.String(200, "test")
}

func postGiveAway(c *gin.Context) {
	c.String(200, "test")
}
