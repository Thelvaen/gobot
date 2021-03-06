package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/Thelvaen/auth/models"
	"github.com/Thelvaen/gobot/templates"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type mailVars struct {
	MainChannel string
	BaseURL     string
	User        string
	UserID      string
	Token       string
}

type tokenStruct struct {
	Token string `json:"password"`
}

func sendMail(ctx iris.Context, token string, id uuid.UUID) {
	var err error
	var user models.User
	var from string

	err = dataStore.Where("id = ?", id).First(&user).Error

	mailT, _ := templates.Asset("mailTemplate.tmpl")
	tmpl := template.Must(template.New("mailTemplate").Parse(string(mailT)))
	if err != nil {
		log.Println("initializing template:", err)
	}

	mail := mailVars{
		MainChannel: conf.Cred.Channel,
		BaseURL:     baseURL(ctx),
		User:        user.Username,
		UserID:      user.ID.String(),
		Token:       token,
	}

	// Sender data.
	if conf.MailConf.From != "" {
		from = conf.MailConf.From
	} else {
		from = "no-reply@twitchbot.domain"
	}

	// Receiver email address.
	to := []string{user.Email}

	// Message body
	message := new(bytes.Buffer)

	err = tmpl.Execute(message, mail)
	if err != nil {
		log.Println("executing template:", err)
	}
	msg := []byte("To:" + strings.Join(to, ",") + "\r\n" + "Subject: Password Reset!\r\n" + "\r\n")
	msg = append(msg, message.Bytes()...)

	var auth smtp.Auth
	if conf.MailConf.Username == "" && conf.MailConf.Password == "" {
		auth = nil
	} else {
		auth = smtp.PlainAuth("", conf.MailConf.Username, conf.MailConf.Password, conf.MailConf.Host)
	}

	err = smtp.SendMail(conf.MailConf.Host+":"+conf.MailConf.Port, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
	}
}
