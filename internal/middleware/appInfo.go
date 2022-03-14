package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("app_name", "blog-service")
		context.Set("app_version", "1.0.0")
		context.Next()
	}
}
