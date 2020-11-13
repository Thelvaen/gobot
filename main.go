package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"gobot/aggregator"
	"gobot/config"
	"gobot/dice"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/gempir/go-twitch-irc/v2"
)

var (
	err      error
	mainChan string
	channels []string

	// Filters gives RegEx and function to call when matching
	filters config.CommandFilter
	// WebRoutes gives endpoints and function to call
	webRoutes config.WebRoutes

	url  string
	port string
)

func init() {
	mainChan = config.BotConfig.Cred.Channel

	filters = make(config.CommandFilter)
	webRoutes = make(config.WebRoutes)

	// Connecting to Twitch
	if config.BotConfig.Cred.IsAuth {
		config.BotConfig.TwitchC = twitch.NewClient(mainChan, config.BotConfig.Cred.Token)
	} else {
		// No credentials provided, anon connection
		config.BotConfig.TwitchC = twitch.NewAnonymousClient()
	}

	// Initializing Web Server for /
	http.HandleFunc("/", getPage)

	// Adding filters & endpoints
	dice.Initialize()
	for filter, modFunction := range dice.Filters {
		filters[filter] = modFunction
	}
	for route, modFunction := range dice.WebRoutes {
		webRoutes[route] = modFunction
	}

	aggregator.Initialize()
	for filter, modFunction := range aggregator.Filters {
		filters[filter] = modFunction
	}
	for route, routeDetails := range aggregator.WebRoutes {
		webRoutes[route] = config.WebTarget{
			RouteFunc: routeDetails.RouteFunc,
			RouteDesc: routeDetails.RouteDesc,
		}
		http.HandleFunc(route, getPage)
	}
}

func pushAndSay(data string) {
	config.BotConfig.TwitchC.Say(mainChan, data)
}

func getPage(w http.ResponseWriter, req *http.Request) {
	for route, routeDetails := range webRoutes {
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
	<a class="navbar-brand" href="/">Fonctions du Bot</a>
	<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
	</button>
	<div class="collapse navbar-collapse" id="navbarNav">
		<ul class="navbar-nav">
`)
	for route, routeDetails := range webRoutes {
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
	if config.BotConfig.Cred.IsAuth {
		if message.Channel == mainChan {
			// Command to process
			for filter, modFunction := range filters {
				found, _ := regexp.MatchString(filter, message.Message)
				if found {
					pushAndSay(modFunction(message))
				}
			}
		}
	}
}

func main() {
	// Starting web server as a Go Routine (background thread)
	go http.ListenAndServe(config.BotConfig.BotServer.URL+":"+strconv.Itoa(config.BotConfig.BotServer.Port), nil)

	config.BotConfig.TwitchC.OnPrivateMessage(func(message twitch.PrivateMessage) {
		parseMessage(message)
	})

	config.BotConfig.TwitchC.Join(mainChan)

	err = config.BotConfig.TwitchC.Connect()
	if err != nil {
		panic(fmt.Errorf("can't connect to Twitch: %s", err))
	}
}
