package config

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/spf13/viper"

	bolt "go.etcd.io/bbolt"
)

var (
	err error
	// BotConfig contains the configuration
	BotConfig Configuration
)

// WebTarget defines WebFunction & Description
type WebTarget struct {
	// RouteFunc gives the module function to be called
	RouteFunc func(*http.Request) string
	// RouteDesc gives the route description
	RouteDesc string
}

// CommandFilter is the map used to store regex & function to parse them
type CommandFilter map[string](func(twitch.PrivateMessage) string)

// WebRoutes is the map used to store routes & function to process them
type WebRoutes map[string]WebTarget

// Configuration object
type Configuration struct {
	// Cred stores credentials
	Cred Credentials
	// Aggreg stores aggregation parameters
	Aggreg Aggregation
	// BotServer stores Webserver parameters
	BotServer WebServer
	// Twitch Client store the Twitch Client interface
	TwitchC *twitch.Client
	// DataStore store the Database
	DataStore *bolt.DB
}

// Credentials define Credential struct
type Credentials struct {
	// IsAuth is set to true if credentials are provided
	IsAuth bool
	// Channel is the IRC Channel/Nickname of the bot
	Channel string
	// Token is the IRC Auth Token
	Token string
}

// Aggregation type to define Aggregation struct
type Aggregation struct {
	// Channels list of chan to aggregate
	Channels []string
	// StackSize
	StackSize int
}

// WebServer defines struct for WebServer parameters
type WebServer struct {
	// Port self
	Port int
	// URL self
	URL string
}

func init() {
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

func main() {

}
