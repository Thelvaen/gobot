package main

import (
	"github.com/Thelvaen/gobot/config"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/kataras/iris/v12/sessions"
	"gorm.io/gorm"
)

type filter struct {
	filterFunc  func(twitch.PrivateMessage) string
	filterRegEx string
}

var (
	sessionsManager *sessions.Sessions
	twitchC         *twitch.Client
	dataStore       *gorm.DB
	filters         []filter
	conf            *config.Configuration
)
