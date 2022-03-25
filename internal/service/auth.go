package service

import (
	"Goose/global"
	"Goose/pkg/email"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//CheckAuth 确认Auth是否存在
func (svc *Service) CheckAuth(authName, authCode string) error {
	auth, err := svc.dao.GetAuth(authName, authCode)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}

//CheckEmail 确认邮箱是否存在
func (svc *Service) CheckEmail(email string) (bool, error) {
	auth, err := svc.dao.GetAuthByEmail(email)

	if err != nil {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

//GenerateCheckCode 生成验证码
func (svc *Service) GenerateCheckCode(mail string) error {
	// 生成验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	checkCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	// 存入Redis
	err := svc.pool.GenerateCheckCode(mail, checkCode)
	if err != nil {
		return err
	}

	// 发送邮箱
	var defaultMailer = email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	content := fmt.Sprintf(`
	<div>
		<div>
			谷声的新朋友，你好呀~
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交了邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, time.Now().Format("2006-01-02 15:04:05"), checkCode)
	err = defaultMailer.SendMailToOne(mail, "谷声账号注册验证", content)
	if err != nil {
		return err
	}

	return nil
}
