package main

import (
	"fmt"
	"net/http"

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
		RouteFunc:     getMessages,
		RouteTemplate: "aggregator.html",
		RouteDesc:     "Aggregateur",
	}

	for _, channel := range BotConfig.Aggreg.Channels {
		BotConfig.TwitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	data := fmt.Sprintf("#%s [%02d:%02d:%02d] <%s> %s", message.Channel, message.Time.Hour(), message.Time.Minute(), message.Time.Second(), message.User.Name, message.Message)
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

func getMessages(req *http.Request) map[string][]string {
	data := map[string][]string{
		"Channels": {},
		"Messages": {},
	}
	for _, channel := range BotConfig.Aggreg.Channels {
		data["Channels"] = append(data["Channels"], channel)
	}
	for i := 0; i < position; i++ {
		data["Messages"] = append(data["Messages"], messages[i])
	}
	return data
}
