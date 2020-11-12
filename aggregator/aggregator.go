package aggregator

import (
	"fmt"
	"gobot/config"
	"net/http"

	"github.com/MakeNowJust/heredoc"
	"github.com/gempir/go-twitch-irc/v2"
)

var (
	err       error
	messages  []string
	position  int
	stackSize int
	mainChan  string
	channels  []string
	// Filters gives RegEx and function to call when matching
	Filters map[string](func(twitch.PrivateMessage) string)
	// WebRoutes gives endpoints and function to call
	WebRoutes map[string]config.WebTarget

	aggregatoreConfig config.Configuration
)

func init() {
	Filters = make(map[string](func(twitch.PrivateMessage) string))
	WebRoutes = make(map[string]config.WebTarget)

	Filters[".*"] = pushMessage
	WebRoutes["/messages"] = config.WebTarget{
		RouteFunc: getMessages,
		RouteDesc: "Aggregateur",
	}
}

// Initialize func allows internals to bet setup after config is loaded during main func init
func Initialize() {
	// Nothing here for this mod
	aggregatoreConfig = config.BotConfig

	stackSize = aggregatoreConfig.Aggreg.StackSize
	mainChan = aggregatoreConfig.Cred.Channel
	messages = make([]string, stackSize+10)
	position = 0
	for _, channel := range aggregatoreConfig.Aggreg.Channels {
		config.BotConfig.TwitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	data := fmt.Sprintf("#%s [%02d:%02d:%02d] &lt;%s&gt; %s", message.Channel, message.Time.Hour(), message.Time.Minute(), message.Time.Second(), message.User.Name, message.Message)
	if position >= stackSize {
		messages[position] = data
		for i := 0; i <= position-1; i++ {
			messages[i] = messages[i+1]
		}
	} else {
		messages[position] = data
		position++
	}
	return ""
}

func getMessages(req *http.Request) (body string) {
	reloadScript := heredoc.Doc(`
<script type="text/javascript" language="javascript">
setTimeout(function(){
	window.location.reload(1);
}, 5000);
</script>
	`)
	body = "<h1>"
	for _, channel := range aggregatoreConfig.Aggreg.Channels {
		body += channel + " "
	}
	body += aggregatoreConfig.Cred.Channel + "</h1><ul>"
	for i := 0; i < position; i++ {
		body += "<li>" + messages[i] + "</li>\n"
	}
	body += "</ul>" + reloadScript
	return
}
