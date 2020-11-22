package main

import (
	auth "github.com/Thelvaen/iris-auth-gorm"

	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/static"
	"github.com/Thelvaen/gobot/templates"

	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/gorilla/securecookie"
)

func webBot() *iris.Application {

	// attach a session manager
	sessionsManager = sessions.New(sessions.Config{
		Cookie:                      config.Cred.Channel,
		Encoding:                    securecookie.New(config.WebConf.HashKey, config.WebConf.BlockKey),
		AllowReclaim:                true,
		DisableSubdomainPersistence: true,
	})

	app := iris.New()
	//app.Logger().SetLevel("debug")

	// Adding sessions
	app.Use(sessionsManager.Handler())

	// Adding CSRF Middleware
	app.Use(csrf.Protect(config.WebConf.CSRF, csrf.Secure(false)))

	// Configuring Auth Middleware
	auth.SetDB(dataStore)
	auth.RequireAuthRoute("/login")

	// Setting it to be used by the router
	app.Use(auth.MiddleWare)

	// Adding context Middleware
	app.Use(prepareContext)
	// Adding Navigation Items
	app.Use(getNavigation)

	// Adding templates & layouts
	tmpl := iris.HTML(templates.AssetFile(), ".html").Reload(true)
	tmpl.Layout("layouts/layout.html")
	app.RegisterView(tmpl)

	// Adding UnAuth Routes
	app.Get("/login", loginHandlerForm)
	app.Get("/logout", logoutHandler)
	app.Post("/login", loginHandler)
	app.Get("/", func(ctx iris.Context) {
		ctx.View("home.html")
	})

	// Adding Auth Routes
	app.PartyFunc("/auth", func(users iris.Party) {
		users.Use(auth.MiddleAuth)
		users.Get("/messages", getMessagesPage)
		users.Get("/stats", getStats)
	})
	app.PartyFunc("/json", func(users iris.Party) {
		users.Use(auth.MiddleAuth)
		users.Get("/messages", getMessagesData)
	})

	// Addming Admin Routes
	app.PartyFunc("/admin", func(users iris.Party) {
		users.Use(auth.MiddleAdmin)
		users.Get("/register", getMessagesData)
	})

	// Adding static content UnAuth
	app.HandleDir("/static", static.AssetFile())

	return app
}
