package errorCode

import "net/http"

var (
	ErrorFormatReg = NewError(20000006, http.StatusInternalServerError, "正则对应错误")
)

var (
	UnauthorizedAuthNotExist   = NewError(20000000, http.StatusInternalServerError, "鉴权失败，找不到对应的 AuthName 和 AuthCode")
	UnauthorizedTokenError     = NewError(20000001, http.StatusInternalServerError, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout   = NewError(20000002, http.StatusUnauthorized, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate  = NewError(20000003, http.StatusInternalServerError, "鉴权失败，Token 生成失败")
	ErrorGenerateCheckCodeFail = NewError(20000004, http.StatusInternalServerError, "生成验证码失败")
	ErrorFormatEmail           = NewError(20000005, http.StatusBadRequest, "邮箱格式错误")
	ErrorFormatAuthName        = NewError(20000007, http.StatusBadRequest, "账号格式错误")
	ErrorFormatAuthCode        = NewError(20000008, http.StatusBadRequest, "密码格式错误")
	ErrorNotValidCheckCode     = NewError(20000009, http.StatusBadRequest, "验证码对应错误")
	ErrorCreateNewAuth         = NewError(20000010, http.StatusInternalServerError, "账号名已存在")
	ErrorAuthNameExist         = NewError(20000011, http.StatusBadRequest, "账号名已存在")
	ErrorEmailExist            = NewError(20000012, http.StatusBadRequest, "邮箱已存在")
	ErrorModifyCode            = NewError(20000013, http.StatusInternalServerError, "修改密码失败")
	ErrorAuthNoExist           = NewError(20000014, http.StatusBadRequest, "账户不存在")
)
