package main

import (
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

	for _, channel := range conf.Aggreg.Channels {
		twitchC.Join(channel)
	}
}

func pushMessage(message twitch.PrivateMessage) string {
	if len(messages) > conf.Aggreg.StackSize {
		i := 0
		copy(messages[i:], messages[i+1:])
		//messages[len(messages)-1] = ""
		messages = messages[:len(messages)-1]
	}
	messages = append(messages, message)
	return ""
}

func getMessagesPage(ctx iris.Context) {
	ctx.ViewData("Channels", conf.Aggreg.Channels)
	if err := ctx.View("aggregator.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func getMessagesData(ctx iris.Context) {
	options := iris.JSON{Indent: "", Secure: false}
	if len(messages) == 0 {
		ctx.JSON("", options)
		return
	}
	ctx.JSON(messages, options)
}
