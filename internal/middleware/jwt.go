package middleware

import (
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWT Token校验
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errorCode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errorCode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errorCode.UnauthorizedTokenTimeout
				default:
					ecode = errorCode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errorCode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}