package main

import (
	"fmt"

	"github.com/Thelvaen/gobot/config"
	auth "github.com/Thelvaen/iris-auth-gorm"
	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
)

type route struct {
	Route string
	Desc  string
}

func prepareContext(ctx iris.Context) {
	userName := ""
	if ctx.User() != nil {
		userName, _ = ctx.User().GetUsername()
	}
	ctx.ViewData("MainChannel", config.Cred.Channel)
	ctx.ViewData("UserName", userName)
	ctx.ViewData("BaseURL", baseURL(ctx))
	ctx.ViewData("CSRFToken", csrf.Token(ctx))
	ctx.Next()
}

func debugMiddle(ctx iris.Context) {
	fmt.Println(ctx.Request().URL.Host)
	ctx.Next()
}

func getNavigation(ctx iris.Context) {
	var navigation []route
	isAuth := auth.IsAuth(ctx)
	if isAuth {
		navigation = append(navigation, route{
			Route: "/auth/messages",
			Desc:  "Aggregateur",
		})
		navigation = append(navigation, route{
			Route: "/auth/stats",
			Desc:  "Statistiques",
		})
	}
	if auth.HasRole(ctx, "admin") {
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
	ctx.ViewData("Navigation", navigation)
	ctx.Next()
}

func baseURL(ctx iris.Context) (url string) {
	scheme := "http://"
	if ctx.Request().TLS != nil {
		scheme = "https://"
	}
	url = scheme + ctx.Request().Host
	return
}
