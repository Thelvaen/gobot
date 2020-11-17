package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
// Block var local
)

// Poll structure made exportable to be used with Gorm ORM
type Poll struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string
	Status      bool
}

// Vote struct allow to gives Roles to user
type Vote struct {
	gorm.Model
	User User
	Poll Poll
}

func initPolls() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  registerVote,
		FilterRegEx: "!vote",
	})
	/*WebRoutes["/polls"] = WebTarget{
		RouteFunc: getVoteForm,
		RouteDesc: "Sondages",
	}*/

	BotConfig.DataStore.AutoMigrate(&Poll{}, &Vote{})
}

func registerVote(message twitch.PrivateMessage) (outMessage string) {
	outMessage = "toto"
	return
}

func getPollListForm(c *gin.Context) {
	c.String(200, "test")
}

func postPollList(c *gin.Context) {
	c.String(200, "test")
}

func getPollForm(c *gin.Context) {
	//pollID := c.Param("poll")
	c.String(200, "test")
}

func postPoll(c *gin.Context) {
	//pollID := c.Param("poll")
	c.String(200, "test")
}
