package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/spf13/viper"
)

var (
	err          error
	twitchC      *twitch.Client
	messages     []string
	position     int
	stackSize    int
	mainChan     string
	channels     []string
	tokenDef     bool
	randomSource *rand.Rand

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

	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))

	stackSize = viper.GetInt("StackSize")
	channels = viper.GetStringSlice("AgregChans")
	mainChan = viper.GetString("MainChannel")

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

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func rollDice(userName string, faces int) {
	// Rolling a dice
	twitchC.Say(mainChan, fmt.Sprintf("* Lance un dé à %d faces pour %s et obtient : %d", faces, userName, randomSource.Intn(faces)+1))
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
	fmt.Fprintf(w, mainChan)
	fmt.Fprintf(w, "</h1>")
	for i := 0; i < position; i++ {
		fmt.Fprintf(w, "<li>%s</li>\n", messages[i])
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func parseMessage(message twitch.PrivateMessage) {
	if (message.Channel == mainChan) && strings.HasPrefix(message.Message, "!") {
		// Command to process
		command := strings.Fields(message.Message)
		switch command[0] {
		case "!dice":
			if len(command) > 1 && isInt(command[1]) {
				// Dice faces
				faces, _ := strconv.Atoi(command[1])
				rollDice(message.User.Name, faces)
			} else {
				rollDice(message.User.Name, 10)
			}
		}
	}
	pushMessage(fmt.Sprintf("#%s [%02d:%02d:%02d] &lt;%s&gt; %s", message.Channel, message.Time.Hour(), message.Time.Minute(), message.Time.Second(), message.User.Name, message.Message))

}

func main() {
	// Starting web server as a Go Routine (background thread)
	go http.ListenAndServe(fmt.Sprintf("%s%s", url, port), nil)

	// Connecting to Twitch
	if viper.IsSet("Credential.Nickname") && viper.IsSet("Credential.Token") {
		twitchC = twitch.NewClient(viper.GetString("Credential.Nickname"), viper.GetString("Credential.Token"))
		//fmt.Println("Authenticated connection")
	} else {
		// No credentials provided, anon connection
		twitchC = twitch.NewAnonymousClient()
		//fmt.Println("Anonymous connection")
	}

	twitchC.OnPrivateMessage(func(message twitch.PrivateMessage) {
		parseMessage(message)
	})

	twitchC.Join(mainChan)
	for _, channel := range channels {
		twitchC.Join(channel)
	}

	err = twitchC.Connect()
	if err != nil {
		panic(fmt.Errorf("can't connect to Twitch: %s", err))
	}
}
