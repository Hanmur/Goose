package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (email *Email) SendMail(to []string, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", email.From)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(email.Host, email.Port, email.UserName, email.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: email.IsSSL}
	return dialer.DialAndSend(message)
}
