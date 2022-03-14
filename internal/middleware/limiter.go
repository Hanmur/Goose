package middleware

import (
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"Goose/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(limit limiter.LimitInterface) gin.HandlerFunc {
	return func(context *gin.Context) {
		key := limit.Key(context)
		if bucket, ok := limit.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(context)
				response.ToErrorResponse(errorCode.TooManyRequests)
				context.Abort()
				return
			}
		}

		context.Next()
	}
}
