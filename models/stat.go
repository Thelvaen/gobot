package models

import (
	"gorm.io/gorm"
)

// Stat structure made exportable to be used with Gorm ORM
type Stat struct {
	gorm.Model
	User  string `gorm:"not null;unique"` // Utilisateur unique!
	Score int    `gorm:"not null;"`
}
