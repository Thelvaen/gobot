package models

import (
	"gorm.io/gorm"
)

// Stat structure made exportable to be used with Gorm ORM
type Stat struct {
	gorm.Model
	TwitchUserID uint
	Score        int `gorm:"not null;"`
}
