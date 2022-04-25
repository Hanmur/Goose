package service

import (
	"Goose/global"
	"Goose/pkg/email"
	"Goose/pkg/errorCode"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//CheckAuth 账户密码校验
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

//SendCheck 发送验证码
func (svc *Service) SendCheck(email, prefix string) *errorCode.Error {
	// 验证邮箱格式
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		global.Logger.ErrorF("svc.SendCheck-regexp.Compile err: %v", err)
		return errorCode.ErrorFormatReg
	}
	if !reg.MatchString(email) {
		return errorCode.ErrorFormatEmail
	}

	// 验证码生成
	err = svc.GenerateCheckCode(email, prefix)
	if err != nil {
		global.Logger.ErrorF("svc.GenerateCheckCode err: %v", err)
		return errorCode.ErrorGenerateCheckCodeFail
	}

	return nil
}

//Register 注册
func (svc *Service) Register(authName, authCode, email, checkCode string) *errorCode.Error {
	// 验证邮箱格式
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		global.Logger.ErrorF("svc.SendCheck-regexp.Compile err: %v", err)
		return errorCode.ErrorFormatReg
	}
	if !reg.MatchString(email) {
		return errorCode.ErrorFormatEmail
	}

	// 验证邮箱是否已存在
	err = svc.CheckNoEmail(email)
	if err != nil {
		global.Logger.ErrorF("svc.CheckEmail err: %v", err)
		return errorCode.ErrorEmailExist
	}

	// 验证账号是否已存在
	err = svc.CheckNoAuthName(authName)
	if err != nil {
		global.Logger.ErrorF("svc.CheckAuthName err: %v", err)
		return errorCode.ErrorAuthNameExist
	}

	// 账号密码格式校验
	err = svc.CheckAuthNameFormat(authName)
	if err != nil {
		return errorCode.ErrorFormatAuthName
	}

	err = svc.CheckAuthCodeFormat(authCode)
	if err != nil {
		return errorCode.ErrorFormatAuthCode
	}

	// 验证码校验
	err = svc.CheckCheckCode(email, checkCode, "regis")
	if err != nil {
		return errorCode.ErrorNotValidCheckCode
	}

	// 账户注册
	err = svc.CreateAuth(authName, authCode, email)
	if err != nil {
		global.Logger.ErrorF("svc.CreateAuth err: %v", err)
		return errorCode.ErrorCreateNewAuth
	}

	return nil
}

//ModifyCode 修改密码
func (svc *Service) ModifyCode(authName, authCode, newCode string) *errorCode.Error {
	// 验证新密码
	err := svc.CheckAuthCodeFormat(newCode)
	if err != nil {
		return errorCode.ErrorFormatAuthCode
	}

	// 获取账户
	err = svc.CheckAuth(authName, authCode)
	if err != nil {
		global.Logger.ErrorF("svc.CheckAuth err: %v", err)
		return errorCode.UnauthorizedAuthNotExist
	}

	// 修改密码
	err = svc.dao.ModifyCode(authName, newCode)
	if err != nil {
		global.Logger.ErrorF("svc.ModifyCode err: %v", err)
		return errorCode.ErrorModifyCode
	}

	return nil
}

//ResetCode 重置密码
func (svc *Service) ResetCode(email, checkCode string) *errorCode.Error {
	// 验证验证码
	err := svc.CheckCheckCode(email, checkCode, "reset")
	if err != nil {
		return errorCode.ErrorNotValidCheckCode
	}

	// 获取账户
	auth, err := svc.dao.GetAuthByEmail(email)
	if err != nil || auth.ID == 0 {
		return errorCode.ErrorAuthNoExist
	}

	// 重置密码
	err = svc.dao.ModifyCode(auth.AuthName, "Goose_007")
	if err != nil {
		global.Logger.ErrorF("svc.ModifyCode err: %v", err)
		return errorCode.ErrorModifyCode
	}

	return nil
}

//----

//CheckNoEmail 确认邮箱是否不存在
func (svc *Service) CheckNoEmail(email string) error {
	auth, err := svc.dao.GetAuthByEmail(email)

	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return errors.New("email does exist")
	}

	return nil
}

//CheckNoAuthName 确认账户名是否不存在
func (svc *Service) CheckNoAuthName(authName string) error {
	auth, err := svc.dao.GetAuthByAuthName(authName)

	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return errors.New("auth does exist")
	}

	return nil
}

//GenerateCheckCode 生成验证码
func (svc *Service) GenerateCheckCode(mail, prefix string) error {
	// 生成验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	checkCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	// 存入Redis
	key := prefix + "_" + mail
	val := checkCode

	err := svc.pool.GenerateCheckCode(key, val)
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
			谷声的朋友，你好呀~
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

//CheckAuthNameFormat 校验账号正则格式
func (svc *Service) CheckAuthNameFormat(authName string) error {
	if len(authName) < 6 {
		return fmt.Errorf("authName is < 6")
	}

	return nil
}

//CheckAuthCodeFormat 校验密码正则格式
func (svc *Service) CheckAuthCodeFormat(authCode string) error {
	if len(authCode) < 6 {
		return fmt.Errorf("password len is < 6")
	}
	num := `[0-9]{1}`
	aToz := `[a-z]{1}`
	AToZ := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, authCode); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(aToz, authCode); !b || err != nil {
		return fmt.Errorf("password need aToz :%v", err)
	}
	if b, err := regexp.MatchString(AToZ, authCode); !b || err != nil {
		return fmt.Errorf("password need AToZ :%v", err)
	}
	if b, err := regexp.MatchString(symbol, authCode); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	}

	return nil
}

//CheckCheckCode 校验验证码
func (svc *Service) CheckCheckCode(email, checkCode, prefix string) error {
	key := prefix + "_" + email
	val := checkCode

	err := svc.pool.CheckCheckCode(key, val)
	if err != nil {
		return err
	}
	return nil
}

//CreateAuth 校验验证码
func (svc *Service) CreateAuth(authName, authCode, email string) error {
	return svc.dao.CreateAuth(authName, authCode, email)
}
