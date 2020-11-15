package main

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
// Block var local
)

func initGiveAway() {
	Filters["!jeveux"] = CLIFilter{
		FilterFunc: registerGiveAway,
		FilterDesc: "Register your participation to a giveaway",
	}
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
