package main

import (
	"github.com/Thelvaen/gobot/config"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12"
)

var (
	messages []twitch.PrivateMessage
)

func initAggregator() {
	filters = append(filters, filter{
		filterFunc:  pushMessage,
		filterRegEx: ".*",
	})

	for _, channel := range config.Aggreg.Channels {
		twitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	if len(messages) > config.Aggreg.StackSize {
		i := 0
		copy(messages[i:], messages[i+1:])
		//messages[len(messages)-1] = ""
		messages = messages[:len(messages)-1]
	}
	messages = append(messages, message)
	return ""
}

func getMessagesPage(ctx iris.Context) {
	if err := ctx.View("aggregator.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func getMessagesData(ctx iris.Context) {
	options := iris.JSON{Indent: "", Secure: false}
	ctx.JSON(messages, options)
}
