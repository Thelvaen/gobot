package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
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
	session := sessions.Get(ctx)
	session.Delete("userID")
	ctx.SetUser(nil)
	ctx.Redirect("/", iris.StatusFound)
}

func loginHandler(ctx iris.Context) {
	session := sessions.Get(ctx)
	var user models.User

	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/login", iris.StatusFound)
		return
	}
	clearPassword := user.Password
	if err := dataStore.Where("name = ?", user.Name).First(&user).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(clearPassword)); err == nil {
			session.Set("userID", user.ID)
			ctx.Redirect("/", iris.StatusFound)
			return
		}
		ctx.Redirect("/login", iris.StatusFound)
		return
	}

}
