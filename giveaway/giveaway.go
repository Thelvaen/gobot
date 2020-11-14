package giveaway

import (
	"net/http"

	"github.com/thelvaen/gobot/config"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	// Filters gives RegEx and function to call when matching
	Filters config.CommandFilter
	// WebRoutes gives endpoints and function to call
	WebRoutes config.WebRoutes

	err            error
	giveAwayConfig config.Configuration
)

func init() {
	Filters = make(config.CommandFilter)
	WebRoutes = make(config.WebRoutes)

	Filters["!jeveux"] = registerGiveAway
	WebRoutes["/giveaway"] = config.WebTarget{
		RouteFunc: getGiveAwayForm,
		RouteDesc: "GiveAway",
	}
}

// Initialize func allows internals to bet setup after config is loaded during main func init
func Initialize() {
	// Nothing here for this mod
	giveAwayConfig = config.BotConfig

}

func registerGiveAway(message twitch.PrivateMessage) (outMessage string) {
	outMessage = "toto"
	return
}

func getGiveAwayForm(req *http.Request) (body string) {
	body = "toto"
	return
}
