package errorCode

import (
	"fmt"
	"net/http"
)

//Error 错误类
type Error struct {
	code int
	msg string
	details []string
}

//codes 错误码表
var codes = map[int]string{}

//NewError 创建Error实例
func NewError(code int, msg string) *Error {
	_, ok := codes[code]
	if ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	codes[code] = msg
	return &Error{code: code, msg: msg}
}

//Error 错误展示
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

//Code 获取错误码
func (e *Error) Code() int {
	return e.code
}

//Msg 获取错误信息
func (e *Error) Msg() string {
	return e.msg
}

//MsgFormats 按照args格式返回错误信息
func (e *Error) MsgFormats(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

//Details 返回错误细节
func (e *Error) Details() []string {
	return e.details
}

//WithDetails 将细节描述添加并返回新的错误
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

//StatusCode 获取错误对应的HTTP状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}