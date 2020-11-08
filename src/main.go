package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	twitchbot "github.com/gempir/go-twitch-irc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	err          error
	client       *twitchbot.Client
	messages     []string
	position     int
	stackSize    int
	channels     []string
	tokenStruct  *oauth2.Token
	tokenDef     bool
	token        string
	oauth2Config *clientcredentials.Config
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

	// Getting OAuth Token if credentials are provided
	if viper.IsSet("Credential.ID") && viper.IsSet("Credential.Secret") {
		oauth2Config = &clientcredentials.Config{
			ClientID:     viper.GetString("Credential.ID"),
			ClientSecret: viper.GetString("Credential.Secret"),
			TokenURL:     twitch.Endpoint.TokenURL,
		}
		tokenStruct, err = oauth2Config.Token(context.Background())
		token = tokenStruct.AccessToken
		if err != nil {
			log.Fatal(err)
		}
		tokenDef = true
	} else {
		tokenDef = false
	}
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
	if tokenDef {
		// Credentials defined in config, auth connection
		client = twitchbot.NewClient(viper.GetString("Credential.Nickname"), fmt.Sprintf("oauth:%s", token))
	} else {
		// No credentials provided, anon connection
		client = twitchbot.NewAnonymousClient()
	}

	client.OnPrivateMessage(func(message twitchbot.PrivateMessage) {
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
