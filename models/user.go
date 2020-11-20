package models

import (
	"strconv"
	"time"

	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

// User structure made exportable to be used with Gorm ORM
type User struct {
	gorm.Model
	Username     string    `gorm:"not null;unique" form:"username" json:"username,omitempty"` // Utilisateur unique!
	Password     string    `gorm:"not null" form:"password" json:"-"`
	Email        string    `gorm:"not null" json:"email,omitempty"`
	Roles        []string  `json:"roles,omitempty"`
	AuthorizedAt time.Time `json:"authorized_at,omitempty"`
}

var _ context.User = (*User)(nil)

// GetAuthorizedAt returns the exact time the
// client has been authorized for the "first" time.
func (u *User) GetAuthorizedAt() (time.Time, error) {
	return u.AuthorizedAt, nil
}

// GetID returns the ID of the User.
func (u *User) GetID() (string, error) {
	return strconv.Itoa(int(u.ID)), nil
}

// GetUsername returns the name of the User.
func (u *User) GetUsername() (string, error) {
	return u.Username, nil
}

// GetPassword returns the raw password of the User.
func (u *User) GetPassword() (string, error) {
	return u.Password, nil
}

// GetEmail returns the e-mail of (string,error) User.
func (u *User) GetEmail() (string, error) {
	return u.Email, nil
}

// GetRoles returns the specific user's roles.
// Returns with `ErrNotSupported` if the Roles field is not initialized.
func (u *User) GetRoles() ([]string, error) {
	if u.Roles == nil {
		return nil, context.ErrNotSupported
	}

	return u.Roles, nil
}
