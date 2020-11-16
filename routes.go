package main

func initRoutes() {
	server.GET("/login", loginHandlerForm)
	server.POST("/login", loginHandler)
	server.GET("/logout", logoutHandler)
	server.GET("/stats", getStats)

	/*authGroup := server.Group("/auth")
	authGroup.Use(middlewareCheckAuth())
	{
		authGroup.GET("/messages", getMessages)
		authGroup.GET("/stats", getStats)
	}*/
}
