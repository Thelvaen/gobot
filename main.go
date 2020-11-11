package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/MakeNowJust/heredoc/v2"
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
	randomSource *rand.Rand
	routes       map[string]string

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

	routes = make(map[string]string)

	if viper.IsSet("Port") {
		port = fmt.Sprintf(":%d", viper.GetInt("Port"))
	} else {
		port = ":8090"
	}
	url = ""

	// Initializing routes
	addRoute("/messages", "Aggregateur de message", getMessages)
	addRoute("/concours", "Concours", getDraw)
	http.HandleFunc("/", getHome)
}

func addRoute(route, desc string, handler func(http.ResponseWriter, *http.Request)) {
	routes[route] = desc
	http.HandleFunc(route, handler)
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func inArray(needle interface{}, haystack interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}

func rollDice(userName string, faces int) {
	// Rolling a dice
	var message string
	message = "* Lance un dé à " + strconv.Itoa(faces) + " faces pour " + userName + " et obtient : " + strconv.Itoa(randomSource.Intn(faces)+1)
	twitchC.Say(mainChan, message)
	pushMessage(fmt.Sprintf("#%s [%02d:%02d:%02d] &lt;%s&gt; %s", mainChan, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), viper.GetString("Credential.Nickname"), message))
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

func getDraw(w http.ResponseWriter, req *http.Request) {
}

func getMessages(w http.ResponseWriter, req *http.Request) {
	var body string
	reloadScript := heredoc.Doc(`
<script type="text/javascript" language="javascript">
setTimeout(function(){
	window.location.reload(1);
}, 5000);
</script>
	`)
	body = "<h1>"
	for _, channel := range channels {
		body += channel + " "
	}
	body += mainChan + "</h1><ul>"
	for i := 0; i < position; i++ {
		body += "<li>" + messages[i] + "</li>\n"
	}
	body += "</ul>" + reloadScript
	getPage(w, body)
}

func getHome(w http.ResponseWriter, req *http.Request) {
	getPage(w, "<h1>Hello World !</h1>")
}

func getPage(w http.ResponseWriter, body string) {
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
	for route, desc := range routes {
		navigation += "<li class=\"nav-item\">" + fmt.Sprintf("<a href=\"%s\">%s</a>", route, desc) + "</li>"
	}
	navigationFooter := heredoc.Doc(`
		</ul>
	</div>
</nav>
`)
	return navigationHeader + navigation + navigationFooter
}

func parseMessage(message twitch.PrivateMessage) {
	pushMessage(fmt.Sprintf("#%s [%02d:%02d:%02d] &lt;%s&gt; %s", message.Channel, message.Time.Hour(), message.Time.Minute(), message.Time.Second(), message.User.Name, message.Message))
	if (message.Channel == mainChan) && strings.HasPrefix(message.Message, "!") {
		// Command to process
		command := strings.Fields(message.Message)
		switch command[0] {
		case "!dice":
			faces := 10
			if len(command) > 1 {
				if !isInt(command[1]) {
					twitchC.Say(mainChan, fmt.Sprintf("J'ai beau essayer, ça je ne vois absolument pas comment faire sans casser toutes les lois de la physique"))
					break
				}
				// Dice faces
				faces, _ = strconv.Atoi(command[1])
				des := []int{2, 3, 4, 6, 8, 10, 12, 16, 20, 24, 100}
				if !inArray(faces, des) {
					break
				}
			}
			rollDice(message.User.Name, faces)
		case "!concours":

		}
	}
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
