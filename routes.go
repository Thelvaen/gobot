package main

import (
	"github.com/Thelvaen/iris-auth-gorm"

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

	// Adding auth Middleware
	auth.SetDB(dataStore)
	auth.RequireAuthRoute("/login")
	app.Use(auth.MiddleWare)

	// Adding context Middleware
	app.Use(prepareContext)

	// Adding templates & layouts
	tmpl := iris.HTML(templates.AssetFile(), ".html").Reload(true)
	tmpl.Layout("layouts/layout.html")
	app.RegisterView(tmpl)

	app.Get("/login", loginHandlerForm)
	app.Get("/logout", logoutHandler)
	app.Post("/login", loginHandler)
	app.Get("/", func(ctx iris.Context) {
		ctx.View("home.html")
	})

	// Adding routes
	app.Get("/auth/messages", getMessagesPage)
	app.Get("/json/messages", getMessagesData)

	app.Get("/auth/stats", getStats)

	// Adding static content
	app.HandleDir("/static", static.AssetFile())

	return app
}
