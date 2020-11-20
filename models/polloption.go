package models

import (
	"gorm.io/gorm"
)

// PollOption struct allow to gives Roles to user
type PollOption struct {
	gorm.Model
	Poll        Poll `gorm:"foreignKey:ID"`
	Name        string
	Description string
	Users       []TwitchUser
}
