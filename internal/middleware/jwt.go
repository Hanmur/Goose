package middleware

import (
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWT Token校验
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			token string
			ecode = errorCode.Success
			auth  *app.Claims
			err   error
		)
		if s, exist := context.GetQuery("token"); exist {
			token = s
		} else {
			token = context.GetHeader("token")
		}
		if token == "" {
			ecode = errorCode.ParamsTokenError
		} else {
			auth, err = app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errorCode.UnauthorizedTokenTimeout
				default:
					ecode = errorCode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errorCode.Success || auth == nil {
			response := app.NewResponse(context)
			response.ToErrorResponse(ecode)
			context.Abort()
			return
		}

		context.Set("auth_name", auth.AuthName)
		context.Next()
	}
}
