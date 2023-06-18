package mailplugin

import mailEngine "github.com/leductoan3082004/go-sdk/plugin/mailer/mail"

type MailEngine interface {
	SendMail(mail mailEngine.Mail) error
}
