package main

func initRoutes() {
	server.GET("/login", loginHandlerForm)
	server.POST("/login", loginHandler)
	server.GET("/logout", logoutHandler)

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
		adminGroup.GET("/registerUser", getNewUserToken)
	}
}
