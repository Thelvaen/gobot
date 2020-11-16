package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thelvaen/gobot/csrf"
)

// Users structure made exportable to be used with Gorm ORM
type Users struct {
	ID       int    `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"not null;unique" form:"username"` // Utilisateur unique!
	Password string `gorm:"not null" form:"password"`
	Email    string `gorm:"not null"`
}

func initAuth() {
	// Initialisation du cookie store

	if !BotConfig.DataStore.HasTable(&Users{}) {
		BotConfig.DataStore.CreateTable(&Users{})
	}

}

func loginHandlerForm(c *gin.Context) {
	c.HTML(200, "login_form.html", gin.H{
		"CSRF_Token": csrf.GetToken(c),
	})
}

func logoutHandler(c *gin.Context) {
	c.Redirect(302, "/")
}

func loginHandler(c *gin.Context) {
	var user Users
	session := sessions.Default(c)

	if c.Bind(&user) == nil {
		hash, _ := hashPassword(user.Password)
		if err := BotConfig.DataStore.Where("name = ? AND password = ?", user.Name, hash).First(&user).Error; err == nil {
			session.Set("userID", user.ID)
			if err := session.Save(); err != nil {
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{
					"error": "Failed to save session",
				})
				return
			}
			c.Redirect(302, "/")
		}
	}

}
