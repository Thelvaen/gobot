package main

import (
	"github.com/Thelvaen/auth"
	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/static"
	"github.com/Thelvaen/gobot/templates"

	"github.com/Thelvaen/csrf"
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
	app.Use(csrf.Protect(config.WebConf.CSRF, csrf.Secure(config.WebConf.IsSecure)))

	// Setting it to be used by the router
	app.Use(auth.Init(auth.Config{
		DataStore:     dataStore,
		LoginRoute:    "/login",
		ReturnOnError: true,
	}))

	// Adding context Middleware
	app.Use(prepareContext)
	// Adding Navigation Items
	app.Use(getNavigation)

	// Adding templates & layouts
	tmpl := iris.HTML(templates.AssetFile(), ".html").Reload(true)
	tmpl.Layout("layouts/layout.html")
	app.RegisterView(tmpl)

	// Adding static content UnAuth
	app.HandleDir("/static", static.AssetFile())

	// Adding UnAuth Routes
	app.Get("/", getHome)
	app.Get("/login", loginHandlerForm)
	app.Get("/logout", logoutHandler)
	app.Post("/login", loginHandler)

	// Adding Auth Routes
	app.PartyFunc("/auth", func(users iris.Party) {
		users.Use(auth.MiddleAuth)
		users.Get("/messages", getMessagesPage)
		users.Get("/stats", getStats)
	})

	app.PartyFunc("/json", func(json iris.Party) {
		json.Use(auth.MiddleAuth)
		json.Get("/messages", getMessagesData)
	})

	// Addming Admin Routes
	app.PartyFunc("/admin", func(admin iris.Party) {
		admin.Use(auth.MiddleRole("admin"))
		admin.Get("/registerUser", createUserForm)
		admin.Post("/registerUser", createUser)
	})

	return app
}

func getHome(ctx iris.Context) {
	if err := ctx.View("home.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}
