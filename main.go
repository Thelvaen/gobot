package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
)

var (
	err error
)

func init() {
	WebRoutes = make(WebRoutesT)
	Filters = make(FiltersT)

	BotConfig.Cred.IsAuth = false
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if !viper.IsSet("Twitch.Channel") {
		panic(fmt.Errorf("variable Twitch Channel must be defined in configuration"))
	} else {
		BotConfig.Cred.Channel = viper.GetString("Twitch.Channel")
	}
	if viper.IsSet("Twitch.Token") {
		BotConfig.Cred.IsAuth = true
		BotConfig.Cred.Token = viper.GetString("Twitch.Token")
	}

	if viper.IsSet("Aggreg.StackSize") {
		BotConfig.Aggreg.StackSize = viper.GetInt("Aggreg.StackSize")
	} else {
		BotConfig.Aggreg.StackSize = 60
	}
	if viper.IsSet("Aggreg.Channels") {
		BotConfig.Aggreg.Channels = viper.GetStringSlice("Aggreg.Channels")
	}

	if viper.IsSet("Http.Port") {
		BotConfig.BotServer.Port = viper.GetInt("Http.Port")
	} else {
		BotConfig.BotServer.Port = 8090
	}
	if viper.IsSet("Http.URL") {
		BotConfig.BotServer.URL = viper.GetString("Http.URL")
	}
	// Opening DB
	BotConfig.DataStore, err = bolt.Open("twitchbot.db", 0600, nil)
	if err != nil {
		panic(fmt.Errorf("can't open BoltDB: %s", err))
	}

	// Intercepting Ctrl+C to close DB properly
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		BotConfig.DataStore.Close()
		os.Exit(0)
	}()
}

func getHome(req *http.Request) (body string) {
	body = "<h1>Hello World !"
	return
}

func pushAndSay(data string) {
	BotConfig.TwitchC.Say(BotConfig.Cred.Channel, data)
}

func getPage(w http.ResponseWriter, req *http.Request) {
	for route, routeDetails := range WebRoutes {
		if req.RequestURI == route {
			servePage(w, routeDetails.RouteFunc(req))
			break
		}
	}
}

func servePage(w http.ResponseWriter, body string) {
	header := heredoc.Doc(`
<!DOCTYPE html>
<html><head><title>Twitch bot</title>
<!-- CSS -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">

<!-- jQuery and JS bundle w/ Popper.js -->
<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js" integrity="sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj" crossorigin="anonymous"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx" crossorigin="anonymous"></script>
`)
	fmt.Fprintf(w, header)
	fmt.Fprintf(w, getNavigation())
	fmt.Fprintf(w, body)
	fmt.Fprintf(w, "</body></html>")
}

func getNavigation() string {
	var navigation string
	navigationHeader := heredoc.Doc(`
<nav class="navbar navbar-expand-lg navbar-light bg-light">
	<a class="navbar-brand" href="#">Fonctions du Bot</a>
	<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
	</button>
	<div class="collapse navbar-collapse" id="navbarNav">
		<ul class="navbar-nav">
`)
	for route, routeDetails := range WebRoutes {
		navigation += "<li class=\"nav-item\">" + fmt.Sprintf("<a href=\"%s\">%s</a>", route, routeDetails.RouteDesc) + "</li>"
	}
	navigationFooter := heredoc.Doc(`
		</ul>
	</div>
</nav>
`)
	return navigationHeader + navigation + navigationFooter
}

func parseMessage(message twitch.PrivateMessage) {
	if BotConfig.Cred.IsAuth {
		// Command to process
		for filter, filterDetails := range Filters {
			found, _ := regexp.MatchString(filter, message.Message)
			if found {
				botProcess := filterDetails.FilterFunc(message)
				if message.Channel == BotConfig.Cred.Channel {
					pushAndSay(botProcess)
				}
			}
		}
	}
}

func initRoutes() {
	for routePath := range WebRoutes {
		http.HandleFunc(routePath, getPage)
	}
}

func main() {
	// Connecting to Twitch
	if BotConfig.Cred.IsAuth {
		BotConfig.TwitchC = twitch.NewClient(BotConfig.Cred.Channel, BotConfig.Cred.Token)
	} else {
		// No credentials provided, anon connection
		BotConfig.TwitchC = twitch.NewAnonymousClient()
	}

	// Registering Twitch IRC Client callback functions
	BotConfig.TwitchC.OnPrivateMessage(func(message twitch.PrivateMessage) {
		parseMessage(message)
	})

	// Initializing modules needs to be done after TwitchConnect
	initAggregator()
	initDice()
	initGiveAway()
	initPolls()

	initRoutes()
	// Starting web server as a Go Routine (background thread)
	url := BotConfig.BotServer.URL + ":" + strconv.Itoa(BotConfig.BotServer.Port)
	go http.ListenAndServe(url, nil)

	BotConfig.TwitchC.Join(BotConfig.Cred.Channel)

	BotConfig.TwitchC.Connect()
	if err != nil {
		panic(fmt.Errorf("can't connect to Twitch: %s", err))
	}
}
