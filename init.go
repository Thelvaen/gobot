package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func init() {
	WebRoutes = make(WebRoutesT)
	//Filters = make(FiltersT)

	BotConfig.Cred.IsAuth = false
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		myPanic("fatal error config file: %s", err)
	}
	if !viper.IsSet("Twitch.Channel") {
		myPanic("variable Twitch Channel must be defined in configuration", nil)
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
		BotConfig.Aggreg.StackSize = 40
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
	BotConfig.DataStore, err = gorm.Open("sqlite3", "twitchbot.db")
	//BotConfig.DataStore.LogMode(true)
	if err != nil {
		myPanic("can't open Sqlite3 DB : %s", err)
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
