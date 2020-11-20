package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/thelvaen/gobot/config"
	"github.com/thelvaen/gobot/models"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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
	dataStore, err = gorm.Open(sqlite.Open("twitchbot.db"), &gorm.Config{})
	//dataStore.LogMode(true)
	if err != nil {
		log.Fatalf("can't open Sqlite3 DB : %s", err)
	}

	seedRoles := true
	seedUsers := true
	if dataStore.Migrator().HasTable("roles") {
		seedRoles = false
	}
	if dataStore.Migrator().HasTable("users") {
		seedUsers = false
	}

	dataStore.Migrator().AutoMigrate(&models.User{}, &models.Role{}, &models.GiveAway{}, &models.Poll{}, &models.PollOption{}, &models.Stat{}, &models.Token{})

	if seedRoles {
		adminRole := models.Role{Name: "admin"}
		userRole := models.Role{Name: "user"}
		dataStore.Create(&adminRole)
		dataStore.Create(&userRole)
	}
	if seedUsers {
		viper.SetConfigName("init")
		viper.MergeInConfig()
		var (
			role      models.Role
			adminRole []models.Role
		)
		err = dataStore.Where("Name = ?", "admin").First(&role).Error
		if err != nil {
			log.Fatalln("can't fetch default admin while seeding the base role :", err)
		}
		adminRole = append(adminRole, role)
		clearPassword := viper.GetString("Seed.Password")
		hashPassword, _ := bcrypt.GenerateFromPassword(
			[]byte(clearPassword),
			bcrypt.DefaultCost,
		)
		userSeed := models.User{
			Name:     viper.GetString("Seed.Name"),
			Password: string(hashPassword),
			Email:    viper.GetString("Seed.Email"),
			Roles:    adminRole,
		}
		dataStore.Create(&userSeed)
	}

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
