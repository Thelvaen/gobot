package main

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/jinzhu/gorm"
)

var (
	// BotConfig contains the configuration
	BotConfig Configuration
	// Filters gives RegEx and function to call when matching
	Filters FiltersT
)

// CLIFilter defines a filter applied to IRC Chat and function to be called
type CLIFilter struct {
	// FilterFunc gives the module function to be called
	FilterFunc func(twitch.PrivateMessage) string
	// FilterRegEx gives the RegEx to apply
	FilterRegEx string
}

// FiltersT is the map used to store regex & function to parse them
type FiltersT []CLIFilter

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
	DataStore *gorm.DB
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
