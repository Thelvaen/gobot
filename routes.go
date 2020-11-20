package main

import (
	"github.com/thelvaen/gobot/config"
	"github.com/thelvaen/gobot/static"
	"github.com/thelvaen/gobot/templates"

	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/gorilla/securecookie"
)

func webBot() *iris.Application {

	// attach a session manager
	cookieName := config.Cred.Channel
	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)
	secureCookie := securecookie.New(hashKey, blockKey)

	sessionsManager = sessions.New(sessions.Config{
		Cookie:       cookieName,
		Encoding:     secureCookie,
		AllowReclaim: true,
	})

	app := iris.New()
	//app.Logger().SetLevel("debug")

	// Adding sessions
	app.Use(sessionsManager.Handler())

	// Adding CSRF Middleware
	app.Use(csrf.Protect(config.WebConf.CSRF, csrf.Secure(false)))

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
