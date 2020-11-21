package models

import "gorm.io/gorm"

// TwitchUser will holds Twitch Users reccords
type TwitchUser struct {
	gorm.Model
	TwitchID    string `gorm:"uniqueIndex;"`
	Name        string
	DisplayName string
	Statistique Stat
	Votes       []PollOption `gorm:"many2many:user_votes;"`
	GiveAways   []GiveAway   `gorm:"many2many:user_giveaways;"`
}
