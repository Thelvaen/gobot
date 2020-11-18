package main

import "github.com/gin-gonic/gin"

func initRoutes() {
	server.GET("/login", loginHandlerForm)
	server.POST("/login", loginHandler)
	server.GET("/logout", logoutHandler)
	server.GET("/", getHome)

	authGroup := server.Group("/auth")
	authGroup.Use(checkAuth())
	{
		authGroup.GET("/messages", getMessagesForm)
		authGroup.GET("/stats", getStats)
	}
	jsonGroup := server.Group("/json")
	jsonGroup.Use(checkAuth())
	{
		jsonGroup.GET("/messages", getMessagesData)
	}
	adminGroup := server.Group("/admin")
	adminGroup.Use(checkAdmin())
	{
		adminGroup.GET("/giveaway", getGiveAwayListForm)
		adminGroup.POST("/giveaway", postGiveAwayList)
		adminGroup.GET("/giveaway/:giveaway", getGiveAwayForm)
		adminGroup.POST("/giveaway/:giveaway", postGiveAway)
		adminGroup.GET("/poll", getPollListForm)
		adminGroup.POST("/poll", postPollList)
		adminGroup.GET("/poll/:poll", getPollForm)
		adminGroup.POST("/poll/:poll", postPoll)
		adminGroup.GET("/registerUser", getNewUserToken)
	}
}

/*
type route struct {
	route string
	desc  string
}*/

func getNavigation(c *gin.Context) (navigation []route) {
	if isAuth(c) {
		navigation = append(navigation, route{
			Route: "/auth/messages",
			Desc:  "Aggregateur",
		})
		navigation = append(navigation, route{
			Route: "/auth/stats",
			Desc:  "Statistiques",
		})
	}
	if isAdmin(c) {
		navigation = append(navigation, route{
			Route: "/admin/giveaway",
			Desc:  "GiveAways",
		})
		navigation = append(navigation, route{
			Route: "/admin/poll",
			Desc:  "Sondages",
		})
		navigation = append(navigation, route{
			Route: "/admin/registerUser",
			Desc:  "Créer un Utilisateur",
		})
	}
	if isAuth(c) {
		navigation = append(navigation, route{
			Route: "/logout",
			Desc:  "Logout",
		})
	} else {
		navigation = append(navigation, route{
			Route: "/login",
			Desc:  "Login",
		})
	}
	return
}
