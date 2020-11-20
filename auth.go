package main

import (
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/thelvaen/gobot/models"
	"golang.org/x/crypto/bcrypt"
)

func loginHandlerForm(ctx iris.Context) {
	if err := ctx.View("login_form.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func logoutHandler(ctx iris.Context) {
	ctx.SetCookieKV("userID", "", iris.CookieAllowSubdomains())
	ctx.Redirect("/", iris.StatusFound)
}

func loginHandler(ctx iris.Context) {
	var user models.User

	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/login", iris.StatusFound)
		return
	}
	clearPassword := user.Password
	if err := dataStore.Where("name = ?", user.Name).First(&user).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(clearPassword)); err == nil {
			ctx.SetCookieKV("userID", strconv.Itoa(int(user.ID)), iris.CookieAllowSubdomains())
			ctx.Redirect("/", iris.StatusFound)
			return
		}
		ctx.Redirect("/login", iris.StatusFound)
		return
	}

}

func isAuth(ctx iris.Context) bool {
	return true
}

func isAdmin(ctx iris.Context) bool {
	return true
}
