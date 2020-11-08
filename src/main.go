package main

import (
	"fmt"
	"net/http"

	"github.com/gempir/go-twitch-irc"
	"github.com/spf13/viper"
)

type config struct {
	Nickname string   `mapstructure:"Nickname"`
	Token    string   `mapstructure:"Token"`
	Channels []string `mapstructure:"Channels"`
	Pile     int      `mapstructure:"Pile"`
}

var (
	// Configuration variable pour le marshalling
	Configuration config
	err           error
	client        *twitch.Client
	messages      []string
	position      int
	taille        int
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = viper.Unmarshal(&Configuration)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	messages = make([]string, Configuration.Pile+10)
	position = 0

	// Initialisation des routes
	http.HandleFunc("/messages", getMessages)
}

func pushMessage(data string) {
	if position >= Configuration.Pile {
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
	// Lancement du serveur Web en routine
	go http.ListenAndServe(":8090", nil)

	// Connection a Twitch
	client = twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		pushMessage(fmt.Sprintf("#%s &lt;%s&gt; %s", message.Channel, message.User.Name, message.Message))
	})

	for _, channel := range Configuration.Channels {
		client.Join(channel)
	}

	err = client.Connect()
	if err != nil {
		panic(err)
	}
}
