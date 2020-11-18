package models

import (
	"gorm.io/gorm"
)

// PollOption struct allow to gives Roles to user
type PollOption struct {
	gorm.Model
	Poll        Poll
	Name        string
	Description string
	Users       []User
}
