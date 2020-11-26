package main

import (
	"encoding/json"

	"github.com/Thelvaen/auth"
	"github.com/Thelvaen/auth/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"golang.org/x/crypto/bcrypt"
)

func loginHandlerForm(ctx iris.Context) {
	if err := ctx.View("loginForm.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func logoutHandler(ctx iris.Context) {
	session := sessions.Get(ctx)
	session.Delete("userID")
	ctx.SetUser(nil)
	ctx.Redirect("/", iris.StatusTemporaryRedirect)
}

func loginHandler(ctx iris.Context) {
	var user models.User

	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	auth.Check(user, ctx)
}

func resetPwdForm(ctx iris.Context) {
	if !ctx.URLParamExists("Token") || !ctx.URLParamExists("UserID") {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	ctx.ViewData("Token", ctx.URLParam("Token"))
	ctx.ViewData("UserID", ctx.URLParam("UserID"))
	if err := ctx.View("resetPwd.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

type sToken struct {
	Token string `json:"password"`
}

func resetPwd(ctx iris.Context) {
	var user models.User
	DBToken := sToken{}
	token := ctx.PostValueDefault("reset.token", "")
	uuid := ctx.PostValueDefault("reset.uuid", "")
	password := ctx.PostValueDefault("password", "")
	if password == "" {
		ctx.ViewData("Token", token)
		ctx.ViewData("UserID", uuid)
		if err := ctx.View("resetPwd.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef(err.Error())
		}
		return
	}
	if token == "" || uuid == "" {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	if err := dataStore.Where("ID = ?", uuid).First(&user).Error; err != nil {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	if err := json.Unmarshal(user.Token, &DBToken); err != nil {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	if DBToken.Token != token {
		ctx.Redirect("/login", iris.StatusTemporaryRedirect)
		return
	}
	buff, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(buff)
	dataStore.Save(&user)
	ctx.Redirect("/login", iris.StatusTemporaryRedirect)
	return
}

func changePwdForm(ctx iris.Context) {
	if err := ctx.View("changePwd.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Writef(err.Error())
	}
}

func changePwd(ctx iris.Context) {
	var user models.User
	oldPassword := ctx.PostValueDefault("oldpassword", "")
	password := ctx.PostValueDefault("newpassword", "")
	if oldPassword == "" || password == "" {
		uid, _ := ctx.User().GetID()
		if err := dataStore.Where("ID = ?", uid).First(&user).Error; err != nil {
			ctx.Redirect("/logout", iris.StatusTemporaryRedirect)
			return
		}
		user.Password = oldPassword
		if !auth.Check(user, ctx) {
			ctx.Redirect("/logout", iris.StatusTemporaryRedirect)
			return
		}
		buff, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(buff)
		dataStore.Save(&user)
		ctx.Redirect("/", iris.StatusTemporaryRedirect)
	}
}
