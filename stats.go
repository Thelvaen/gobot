package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	_ "github.com/mattn/go-sqlite3"
)

var (
	statement *sql.Stmt
)

func initStats() {
	messages = make([]string, BotConfig.Aggreg.StackSize+10)

	Filters[".*"] = CLIFilter{
		FilterFunc: pushStats,
		FilterDesc: "Get every message to stats",
	}

	Filters["^!score$"] = CLIFilter{
		FilterFunc: getCliStats,
		FilterDesc: "Output messages to the chan",
	}

	WebRoutes["/stats"] = WebTarget{
		RouteFunc:     getStats,
		RouteTemplate: "stats.html",
		RouteDesc:     "Statistiques",
	}

	// Creating Stats bucket
	createStatsTable := `CREATE TABLE IF NOT EXISTS stats (
		id INTEGER PRIMARY KEY,
		user TEXT NOT NULL,
		score INTEGER NOT NULL
		);`
	BotConfig.DataStore.Exec(createStatsTable)
}

func pushStats(message twitch.PrivateMessage) string {
	if message.Channel != BotConfig.Cred.Channel {
		return ""
	}
	if len(message.Message) < 10 {
		return ""
	}
	if strings.HasPrefix(message.Message, "!") {
		return ""
	}

	var currentScore int
	statement, err = BotConfig.DataStore.Prepare("SELECT score FROM stats WHERE user = ?")
	defer statement.Close()
	if err != nil {
		myPanic("error while preparing statement: ", err)
	}
	err = statement.QueryRow(message.User.Name).Scan(&currentScore)
	if err == sql.ErrNoRows {
		statement, err = BotConfig.DataStore.Prepare("INSERT INTO stats(user, score) values(?, ?)")
		if err != nil {
			myPanic("error inserting initial data: ", err)
		}
		statement.Exec(message.User.Name, 0)
		currentScore = 0
	}
	if err != nil && err != sql.ErrNoRows {
		myPanic("error while fetching data: ", err)
	}
	currentScore++
	statement, err = BotConfig.DataStore.Prepare("UPDATE stats SET score = ? WHERE user = ?")
	statement.Exec(currentScore, message.User.Name)

	return ""
}

func getCliStats(message twitch.PrivateMessage) string {
	// Outputting stats to the channel
	if message.Channel != BotConfig.Cred.Channel {
		return ""
	}

	var currentScore int
	statement, err = BotConfig.DataStore.Prepare("SELECT score FROM stats WHERE user = ?")
	defer statement.Close()
	if err != nil {
		myPanic("error while preparing statement: ", err)
	}
	err = statement.QueryRow(message.User.Name).Scan(&currentScore)
	if err == sql.ErrNoRows {
		statement, err = BotConfig.DataStore.Prepare("INSERT INTO stats(user, score) values(?, ?)")
		if err != nil {
			myPanic("error inserting initial data: ", err)
		}
		statement.Exec(message.User.Name, 0)
		currentScore = 0
	}

	return "Ton score est : " + strconv.Itoa(currentScore)
}

func getStats(req *http.Request) map[string]map[string]string {
	var rows *sql.Rows
	data := map[string]map[string]string{
		"Statistiques": {},
	}
	rows, err = BotConfig.DataStore.Query("SELECT user, score FROM stats")
	defer rows.Close()
	for rows.Next() {
		var user string
		var score int
		err = rows.Scan(&user, &score)
		data["Statistiques"][user] = strconv.Itoa(score)
	}
	return data
}
