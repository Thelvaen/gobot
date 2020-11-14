package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	randomSource *rand.Rand
)

func initDice() {
	Filters["^!dice$"] = CLIFilter{
		FilterFunc: RollDice,
		FilterDesc: "Rolls a dice without faces specified",
	}
	Filters["^!dice \\d*$"] = CLIFilter{
		FilterFunc: RollDice,
		FilterDesc: "Rolls a dice faces have been specified",
	}
	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RollDice called by the bot when rolling a dice
func RollDice(message twitch.PrivateMessage) (dice string) {
	// Rolling a dice
	dice = "J'ai beau essayer, ça je ne vois absolument pas comment faire sans casser toutes les lois de la physique"
	faces := 10

	command := strings.Fields(message.Message)

	if len(command) > 1 {
		if (len(command[1]) > 0) && !isInt(command[1]) {
			return
		}
		faces, _ = strconv.Atoi(command[1])
		des := []int{2, 3, 4, 6, 8, 10, 12, 16, 20, 24, 100}
		if !inArray(faces, des) {
			dice = "Essaye de choisir une valeur de dé existante"
			return
		}
	}
	dice = "* Lance un dé à " + strconv.Itoa(faces) + " faces pour " + message.User.Name + " et obtient : " + strconv.Itoa(randomSource.Intn(faces)+1)
	return
}
