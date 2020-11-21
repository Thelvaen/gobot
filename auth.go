package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/thelvaen/iris-auth-gorm"
	"github.com/thelvaen/iris-auth-gorm/models"
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
	var user models.User

	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/login", iris.StatusFound)
		return
	}
	auth.Check(user, ctx)
}
