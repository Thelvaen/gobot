package main

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
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

func getGiveAwayForm(req *http.Request) (body string) {
	body = "toto"
	return
}
