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
	Filters = append(Filters, CLIFilter{
		FilterFunc:  RollDice,
		FilterRegEx: "^!dice$",
	})

	Filters = append(Filters, CLIFilter{
		FilterFunc:  RollDice,
		FilterRegEx: "^!dice \\d*$",
	})

	Filters = append(Filters, CLIFilter{
		FilterFunc:  RollRand,
		FilterRegEx: "^!rand \\d*$",
	})
	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RollDice called by the bot when rolling a dice
func RollDice(message twitch.PrivateMessage) (dice string) {
	// Rolling a dice
	faces := 10

	command := strings.Fields(message.Message)

	if len(command) > 1 {
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

// RollRand called by the bot when rolling a dice
func RollRand(message twitch.PrivateMessage) (dice string) {
	// Rolling a dice
	dice = "J'ai beau essayer, ça je ne vois absolument pas comment faire sans casser toutes les lois de la physique"
	faces := 100

	command := strings.Fields(message.Message)

	if len(command) > 1 {
		if (len(command[1]) > 0) && !isInt(command[1]) {
			return
		}
		faces, _ = strconv.Atoi(command[1])
	}
	dice = "* Choisi un nombre entre 1 et " + strconv.Itoa(faces) + " pour " + message.User.Name + " et obtient : " + strconv.Itoa(randomSource.Intn(faces)+1)
	return
}
