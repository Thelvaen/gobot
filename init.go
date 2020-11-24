package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	var err error
	conf, err = config.NewConfig()
	if err == config.ErrorLoading {
		log.Fatalln("can't open conf.yml file")
	}
	if err == config.ErrorLogin {
		log.Fatalln("main Channel not provided in the config")
	}
	if err != nil {
		log.Fatalln(err)
	}

	// Configuring Gorm & Gorm Logger

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)

	dbConf := gorm.Config{
		FullSaveAssociations: true,
		Logger:               gormLogger,
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
