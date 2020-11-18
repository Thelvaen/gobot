package models

import (
	"gorm.io/gorm"
)

// GiveAway structure made exportable to be used with Gorm ORM
type GiveAway struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string
	Status      bool
	Users       []User
}
