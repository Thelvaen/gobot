package main

import (
	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
)

type route struct {
	Route string
	Desc  string
}

func prepareContext(ctx iris.Context) {
	ctx.ViewData("UserName", getUserName(ctx))
	ctx.ViewData("BaseURL", baseURL(ctx))
	ctx.ViewData("Navigation", getNavigation(ctx))
	ctx.ViewData("CSRFToken", csrf.Token(ctx))
	ctx.Next()
}

func getNavigation(ctx iris.Context) (navigation []route) {
	if isAuth(ctx) {
		navigation = append(navigation, route{
			Route: "/auth/messages",
			Desc:  "Aggregateur",
		})
		navigation = append(navigation, route{
			Route: "/auth/stats",
			Desc:  "Statistiques",
		})
	}
	if isAdmin(ctx) {
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
	if isAuth(ctx) {
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

func baseURL(ctx iris.Context) string {
	return ""
}

func getUserName(ctx iris.Context) string {
	return ""
}
