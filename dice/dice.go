package dice

import (
	"gobot/config"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	// Filters gives RegEx and function to call when matching
	Filters config.CommandFilter
	// WebRoutes gives endpoints and function to call
	WebRoutes config.WebRoutes

	randomSource *rand.Rand
	diceConfig   config.Configuration
)

func init() {
	Filters = make(config.CommandFilter)
	WebRoutes = make(config.WebRoutes)

	Filters["^!dice"] = RollDice
	randomSource = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Initialize func allows internals to bet setup after config is loaded during main func init
func Initialize() {
	// Nothing here for this mod
	diceConfig = config.BotConfig
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
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
		faces, _ := strconv.Atoi(command[1])
		des := []int{2, 3, 4, 6, 8, 10, 12, 16, 20, 24, 100}
		if !inArray(faces, des) {
			dice = "Essaye de choisir une valeur de dé existante"
			return
		}
	}
	dice = "* Lance un dé à " + strconv.Itoa(faces) + " faces pour " + message.User.Name + " et obtient : " + strconv.Itoa(randomSource.Intn(faces)+1)
	return
}
