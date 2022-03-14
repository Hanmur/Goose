package middleware

import (
	"Goose/global"
	"Goose/pkg/app"
	"Goose/pkg/email"
	"Goose/pkg/errorCode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultEmail := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.ErrorF("panic recover err: %v", err)

				err := defaultEmail.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.PanicF("email.SendMail err: %v", err)
				}

				app.NewResponse(context).ToErrorResponse(errorCode.ServerError)
				context.Abort()
			}
		}()
		context.Next()
	}
}
