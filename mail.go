package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"text/template"

	"github.com/Thelvaen/gobot/config"
	"github.com/Thelvaen/gobot/templates"
	"github.com/Thelvaen/iris-auth-gorm/models"
)

type mailVars struct {
	MainChannel string
	BaseURL     string
	User        string
	Token       string
}

type tokenStruct struct {
	Token string `json:"password"`
}

func sendMail(token string, id uint) {
	var err error
	var user models.User
	var from string

	err = dataStore.Debug().Where("id = ?", id).First(&user).Error

	mailT, _ := templates.Asset("mailTemplate.tmpl")
	tmpl := template.Must(template.New("mailTemplate").Parse(string(mailT)))
	if err != nil {
		log.Println("initializing template:", err)
	}

	mail := mailVars{
		MainChannel: config.Cred.Channel,
		BaseURL:     config.WebConf.URL,
		User:        user.Username,
		Token:       token,
	}

	// Sender data.
	if config.MailConf.From != "" {
		from = config.MailConf.From
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
	if config.MailConf.Username == "" && config.MailConf.Password == "" {
		auth = nil
		fmt.Println("auth nil")
	} else {
		auth = smtp.PlainAuth("", config.MailConf.Username, config.MailConf.Password, config.MailConf.Host)
	}

	err = smtp.SendMail(config.MailConf.Host+":"+config.MailConf.Port, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
	}
}
