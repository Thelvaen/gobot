package models

import (
	"gorm.io/gorm"
)

// Token table holds creation token for user registering
type Token struct {
	gorm.Model
	Token string `gorm:"not null;"`
	Users User   `gorm:"foreignKey:ID"`
}
