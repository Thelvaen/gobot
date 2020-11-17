package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
)

var (
// Block var local
)

func initGiveAway() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  registerGiveAway,
		FilterRegEx: "^!jeveux$",
	})
	/*WebRoutes["/giveaway"] = WebTarget{
		RouteFunc: getGiveAwayForm,
		RouteDesc: "GiveAway",
	}*/
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
