package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
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

func getMessages(c *gin.Context) {
	data := map[string]map[string]string{
		"Channels": {},
		"Messages": {},
	}
	i := 0
	for _, channel := range BotConfig.Aggreg.Channels {
		data["Channels"][strconv.Itoa(i)] = channel
		i++
	}
	for i := 0; i < position; i++ {
		data["Messages"][strconv.Itoa(i)] = messages[i]
	}
	c.HTML(http.StatusOK, "aggregator.html", gin.H{
		"MainChannel": BotConfig.Cred.Channel,
		"WebRoutes":   WebRoutes,
		"Data":        data,
	})
}
