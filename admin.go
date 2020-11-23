package main

import (
	auth "github.com/Thelvaen/iris-auth-gorm"
	"github.com/Thelvaen/iris-auth-gorm/models"
	"github.com/kataras/iris/v12"
)

func createUserForm(ctx iris.Context) {
	if err := ctx.View("createUser.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func createUser(ctx iris.Context) {
	var user models.User

	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/admin/registerUser", iris.StatusFound)
		return
	}
	sendMail(auth.CreateUser(user))
}
