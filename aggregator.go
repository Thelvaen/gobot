package main

import (
	"net/http"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gin-gonic/gin"
)

var (
	messages  []twitch.PrivateMessage
	position  int
	stackSize int
)

func initAggregator() {
	Filters = append(Filters, CLIFilter{
		FilterFunc:  pushMessage,
		FilterRegEx: ".*",
	})

	for _, channel := range BotConfig.Aggreg.Channels {
		BotConfig.TwitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	if len(messages) > BotConfig.Aggreg.StackSize {
		i := 0
		copy(messages[i:], messages[i+1:])
		//messages[len(messages)-1] = ""
		messages = messages[:len(messages)-1]
	}
	messages = append(messages, message)
	return ""
}

func getMessagesForm(c *gin.Context) {
	Channels := make(map[string]string)
	//Messages := make(map[string]string)
	i := 0
	for _, channel := range BotConfig.Aggreg.Channels {
		Channels[strconv.Itoa(i)] = channel
		i++
	}
	c.HTML(http.StatusOK, "aggregator.html", gin.H{
		"BaseURL":     baseURL(c),
		"MainChannel": BotConfig.Cred.Channel,
		"WebRoutes":   WebRoutes,
		"Channels":    Channels,
	})
}

func getMessagesData(c *gin.Context) {
	c.JSON(200, messages)
}
