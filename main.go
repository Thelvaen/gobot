package main

import (
	"fmt"
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/spf13/viper"
)

var (
	err       error
	client    *twitch.Client
	messages  []string
	position  int
	stackSize int
	channels  []string
	tokenDef  bool

	url  string
	port string
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

	tokenDef = true

	if viper.IsSet("Port") {
		port = fmt.Sprintf(":%d", viper.GetInt("Port"))
	} else {
		port = ":8090"
	}
	url = ""

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
	fmt.Fprintf(w, "<h1>")
	for _, channel := range channels {
		fmt.Fprintf(w, channel)
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "</h1>")
	for i := 0; i < position; i++ {
		fmt.Fprintf(w, "<li>%s</li>\n", messages[i])
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func main() {
	// Starting web server as a Go Routine (background thread)
	go http.ListenAndServe(fmt.Sprintf("%s%s", url, port), nil)

	// Connecting to Twitch
	if viper.IsSet("Credential.Nickname") && viper.IsSet("Credential.Token") {
		client = twitch.NewClient(viper.GetString("Credential.Nickname"), viper.GetString("Credential.Token"))
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
		panic(fmt.Errorf("can't connect to Twitch: %s", err))
	}
}
