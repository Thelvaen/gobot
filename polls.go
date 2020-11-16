package main

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
// Block var local
)

func initPolls() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  registerVote,
		FilterRegEx: "!vote",
	})
	/*WebRoutes["/polls"] = WebTarget{
		RouteFunc: getVoteForm,
		RouteDesc: "Sondages",
	}*/
}

func registerVote(message twitch.PrivateMessage) (outMessage string) {
	outMessage = "toto"
	return
}

func getVoteForm(req *http.Request) (body string) {
	body = "toto"
	return
}
