package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	err := config.LoadAndParse()
	if err == config.ErrorLoading {
		log.Fatalln("can't open config.yml file")
	}
	if err == config.ErrorLogin {
		log.Fatalln("main Channel not provided in the config")
	}
	if err != nil {
		log.Fatalln(err)
	}

	// Opening DB
	dbConf := gorm.Config{
		FullSaveAssociations: true,
		Logger:               nil,
	}
	dataStore, err = gorm.Open(sqlite.Open("twitchbot.db"), &dbConf)
	if err != nil {
		log.Fatalf("can't open Sqlite3 DB : %s", err)
	}

	dataStore.Migrator().AutoMigrate(&models.TwitchUser{}, &models.GiveAway{}, &models.Poll{}, &models.PollOption{}, &models.Stat{})

	// Intercepting Ctrl+C to close DB properly
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		db, _ := dataStore.DB()
		db.Close()
		os.Exit(0)
	}()
}
