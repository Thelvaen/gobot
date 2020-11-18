package models

import (
	"gorm.io/gorm"
)

// Role struct allow to gives Roles to user
type Role struct {
	gorm.Model
	Name  string `gorm:"not null;unique"` // Role unique!
	Users []User `gorm:"many2many:user_roles;"`
}
