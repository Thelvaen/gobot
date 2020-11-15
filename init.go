package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	bolt "go.etcd.io/bbolt"
)

func init() {
	WebRoutes = make(WebRoutesT)
	Filters = make(FiltersT)

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
		myPanic("can't open BoltDB: %s", err)
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
