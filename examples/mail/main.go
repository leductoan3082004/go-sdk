package main

import (
	"fmt"
	goservice "github.com/leductoan3082004/go-sdk"
	mailplugin "github.com/leductoan3082004/go-sdk/plugin/mailer"
	mailengine "github.com/leductoan3082004/go-sdk/plugin/mailer/mail"
	"log"
)

func main() {
	sc := goservice.New(
		goservice.WithName("mail"),
		goservice.WithVersion("hihi"),
		goservice.WithInitRunnable(mailengine.NewMailEngine("mail", "mail")),
	)
	sc.OutEnv()
	mail := sc.MustGet("mail").(mailplugin.MailEngine)
	data := mailengine.Mail{
		Subject:  "Your verification link for hareta",
		Body:     fmt.Sprintf("Click this link to verify your email: <a href='http://localhost:8080/check-verification-code"),
		Receiver: "thanhthaofirecrush@gmail.com",
	}
	if err := mail.SendMail(data); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("success")
}
