package polls

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/thelvaen/gobot/config"
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

	Filters["!vote"] = registerVote
	WebRoutes["/polls"] = config.WebTarget{
		RouteFunc: getVoteForm,
		RouteDesc: "Sondages",
	}
}

// Initialize func allows internals to bet setup after config is loaded during main func init
func Initialize() {
	// Nothing here for this mod
	giveAwayConfig = config.BotConfig

}

func registerVote(message twitch.PrivateMessage) (outMessage string) {
	outMessage = "toto"
	return
}

func getVoteForm(req *http.Request) (body string) {
	body = "toto"
	return
}
