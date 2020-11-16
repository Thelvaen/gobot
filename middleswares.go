package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CheckAuth middlewares check if user is auth via session
func checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID != nil {
			c.Next()
			return
		}

		c.Redirect(302, "/login")
		c.Next()
		return
	}
}

// CheckAdmin middlewares test if the user is admin or not
func checkAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID != nil {
			var user User
			if err := BotConfig.DataStore.Preload("Roles").Where("ID = ?", userID).First(&user).Error; err == nil {
				for _, role := range user.Roles {
					fmt.Println(role.Name)
					if role.Name == "admin" {
						c.Next()
						return
					}
				}
			}
		}

		c.Redirect(302, "/login")
		c.Next()
		return
	}
}
