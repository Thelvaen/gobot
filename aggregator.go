package main

import (
	"fmt"
	"net/http"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/gempir/go-twitch-irc/v2"
)

var (
	messages  []string
	position  int
	stackSize int
)

func initAggregator() {
	messages = make([]string, BotConfig.Aggreg.StackSize+10)

	Filters[".*"] = CLIFilter{
		FilterFunc: pushMessage,
		FilterDesc: "Get every message to aggregator",
	}

	WebRoutes["/messages"] = WebTarget{
		RouteFunc: getMessages,
		RouteDesc: "Aggregateur",
	}

	for _, channel := range BotConfig.Aggreg.Channels {
		BotConfig.TwitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	data := fmt.Sprintf("#%s [%02d:%02d:%02d] &lt;%s&gt; %s", message.Channel, message.Time.Hour(), message.Time.Minute(), message.Time.Second(), message.User.Name, message.Message)
	if position >= BotConfig.Aggreg.StackSize {
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
	for _, channel := range BotConfig.Aggreg.Channels {
		body += channel + " "
	}
	body += BotConfig.Cred.Channel + "</h1><ul>"
	for i := 0; i < position; i++ {
		body += "<li>" + messages[i] + "</li>\n"
	}
	body += "</ul>" + reloadScript
	return
}
