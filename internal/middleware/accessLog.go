package middleware

import (
	"Goose/global"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (writer AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := writer.body.Write(p); err != nil {
		return n, err
	}
	return writer.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: context.Writer}
		context.Writer = bodyWriter

		beginTime := time.Now().Unix()
		context.Next()
		endTime := time.Now().Unix()

		fields := "request{" + context.Request.PostForm.Encode() + "}, response{" + bodyWriter.body.String() + "}"
		global.Logger.InfoF("fields: %s, access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			fields,
			context.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
