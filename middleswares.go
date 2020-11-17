package main

import (
	"github.com/gin-gonic/gin"
)

// CheckAuth middlewares check if user is auth via session
func checkAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		if isAuth(c) {
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
		if isAdmin(c) {
			c.Next()
			return
		}

		c.Redirect(302, "/login")
		c.Next()
		return
	}
}
