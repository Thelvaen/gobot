package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thelvaen/gobot/csrf"
)

type route struct {
	Route string
	Desc  string
}

type myContext struct {
	UserName    interface{}
	BaseURL     string
	MainChannel string
	CSRFToken   string
	Navigation  []route
}

func prepareContext(c *gin.Context) myContext {
	return myContext{
		UserName:    getUserName(c),
		BaseURL:     baseURL(c),
		CSRFToken:   csrf.GetToken(c),
		Navigation:  getNavigation(c),
		MainChannel: BotConfig.Cred.Channel,
	}
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Context": prepareContext(c),
	})
}
