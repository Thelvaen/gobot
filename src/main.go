package main

import (
	"fmt"
	"net/http"

	"github.com/gempir/go-twitch-irc"
	"github.com/spf13/viper"
)

var (
	err       error
	client    *twitch.Client
	messages  []string
	position  int
	stackSize int
	channels  []string
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	stackSize = viper.GetInt("StackSize")
	channels = viper.GetStringSlice("AgregChans")

	messages = make([]string, stackSize+10)
	position = 0

	// Initializing routes
	http.HandleFunc("/messages", getMessages)
}

func pushMessage(data string) {
	if position >= stackSize {
		messages[position] = data
		for i := 0; i <= position-1; i++ {
			messages[i] = messages[i+1]
		}
	} else {
		messages[position] = data
		position++
	}
}

func getMessages(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><head><title>Aggregateur de message</title><meta content='5' http-equiv='refresh'/></head><body><ul>")
	for i := 0; i < position; i++ {
		fmt.Fprintf(w, "<li>%s</li>\n", messages[i])
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func main() {
	// Starting web server as a Go Routine (background thread)
	go http.ListenAndServe(":8090", nil)

	// Connecting to Twitch
	if viper.IsSet("Credential.Nickname") && viper.IsSet("Credential.Token") {
		// Credentials defined in config, auth connection
	} else {
		// No credentials provided, anon connection
		client = twitch.NewAnonymousClient()
	}

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		pushMessage(fmt.Sprintf("#%s &lt;%s&gt; %s", message.Channel, message.User.Name, message.Message))
	})

	for _, channel := range channels {
		client.Join(channel)
	}

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
