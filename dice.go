package main

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	randomSource *rand.Rand
)

func initDice() {
	filters = append(filters, filter{
		filterFunc:  RollDice,
		filterRegEx: "^!dice$",
	})

	filters = append(filters, filter{
		filterFunc:  RollDice,
		filterRegEx: "^!dice \\d*$",
	})

	filters = append(filters, filter{
		filterFunc:  RollRand,
		filterRegEx: "^!rand \\d*$",
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
	command := strings.Fields(message.Message)

	faces, _ := strconv.Atoi(command[1])

	dice = "* Choisi un nombre entre 1 et " + strconv.Itoa(faces) + " pour " + message.User.Name + " et obtient : " + strconv.Itoa(randomSource.Intn(faces)+1)
	return
}

func inArray(needle interface{}, haystack interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}
