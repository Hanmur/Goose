package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(timeContext *gin.Context) {
	return func(timeContext *gin.Context) {
		ctx, cancel := context.WithTimeout(timeContext.Request.Context(), t)
		defer cancel()

		timeContext.Request = timeContext.Request.WithContext(ctx)
		timeContext.Next()
	}
}
