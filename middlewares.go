package main

import (
	"fmt"

	"github.com/iris-contrib/middleware/csrf"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/thelvaen/gobot/models"
)

type route struct {
	Route string
	Desc  string
}

func startSession(ctx iris.Context) {
	session := sessions.Get(ctx)
	userID := session.Get("userID")
	if userID == nil {
		ctx.Next()
	}
	var user models.User
	if err := dataStore.Preload("Roles").Where("ID = ?", userID).First(&user).Error; err == nil {
		fmt.Println(ctx.SetUser(user))
	}
	fmt.Println(ctx.User())
	ctx.Next()
}

func prepareContext(ctx iris.Context) {
	ctx.ViewData("UserName", (ctx))
	ctx.ViewData("BaseURL", baseURL(ctx))
	ctx.ViewData("CSRFToken", csrf.Token(ctx))
	ctx.Next()
}

func getNavigation(ctx iris.Context) {
	var navigation []route
	if ctx.User() != nil {
		navigation = append(navigation, route{
			Route: "/auth/messages",
			Desc:  "Aggregateur",
		})
		navigation = append(navigation, route{
			Route: "/auth/stats",
			Desc:  "Statistiques",
		})
	}
	/*if isAdmin(ctx) {
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
	}*/
	if ctx.User() != nil {
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
	ctx.ViewData("Navigation", navigation)
	ctx.Next()
}

func baseURL(ctx iris.Context) string {
	return ""
}
