package email

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

//defailtMailer := email.NewEmail(&email.SMTPInfo{
//	Host:     global.EmailSetting.Host,
//	Port:     global.EmailSetting.Port,
//	IsSSL:    global.EmailSetting.IsSSL,
//	UserName: global.EmailSetting.UserName,
//	Password: global.EmailSetting.Password,
//	From:     global.EmailSetting.From,
//})
//var mailTo []string
//mailTo = append(mailTo, "1466046208@qq.com")
//err = defailtMailer.SendMail(mailTo, "Test", "We are trying to send a email.")
//if err != nil{
//	global.Logger.Info("Err send email")
//}
