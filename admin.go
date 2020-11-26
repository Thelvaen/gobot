package main

import (
	auth "github.com/Thelvaen/auth"
	"github.com/Thelvaen/auth/models"
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

	roles := []string{"user"}
	err := ctx.ReadForm(&user)
	user.Roles = roles
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/admin/registerUser", iris.StatusFound)
		return
	}
	token, uuid := auth.CreateUser(user)
	sendMail(ctx, token, uuid)
}
