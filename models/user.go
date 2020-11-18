package models

import (
	"gorm.io/gorm"
)

// User structure made exportable to be used with Gorm ORM
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique" form:"username"` // Utilisateur unique!
	Password string `gorm:"not null" form:"password"`
	Email    string `gorm:"not null"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}
