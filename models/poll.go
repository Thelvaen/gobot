package models

import (
	"gorm.io/gorm"
)

// Poll structure made exportable to be used with Gorm ORM
type Poll struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string
	Status      bool
}
