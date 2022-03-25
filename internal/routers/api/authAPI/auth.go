package authAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/internal/service/validator"
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
	"regexp"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

//Login
// @Summary  	登录
// @Description	登录，获取Token
// @Tags	 	账户管理
// @Produce  	json
// @Param    	auth_name   formData     string   	true  	"认证账号" 	default(Hanmur)
// @Param    	auth_code  	formData     string   	true  	"认证密码" 	default(Hanmur_Goose)
// @Success  	200        {object}  string      		"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/auth/login [POST]
func (auth Auth) Login(c *gin.Context) {
	// 参数校验
	param := validator.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 账户检测
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(param.AuthName, param.AuthCode)
	if err != nil {
		global.Logger.ErrorF("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errorCode.UnauthorizedAuthNotExist)
		return
	}

	// Token生成
	token, err := app.GenerateToken(param.AuthName, param.AuthCode)
	if err != nil {
		global.Logger.ErrorF("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errorCode.UnauthorizedTokenGenerate)
		return
	}

	// 响应
	response.ToResponse(gin.H{
		"token": token,
	})
}

//SendCheck
// @Summary  	发送验证码
// @Description	在Redis中生成验证码并发送该验证码至Redis
// @Tags	 	账户管理
// @Produce  	json
// @Param    	email   	formData     string   	true  	"邮箱" 	default(1466046208@qq.com)
// @Success  	200        {object}  string      		"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/auth/sendCheck [POST]
func (auth Auth) SendCheck(c *gin.Context) {
	// 参数校验
	param := validator.GetEmailRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	// 验证邮箱格式
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		global.Logger.ErrorF("regexp err: ", err)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if !reg.MatchString(param.Email) {
		response.ToResponse("邮箱格式错误")
		return
	}

	// 验证邮箱是否已存在
	isExist, err := svc.CheckEmail(param.Email)
	if err != nil {
		global.Logger.ErrorF("svc.CheckEmail err: %v", err)
		response.ToErrorResponse(errorCode.ErrorSendCheckFail)
		return
	}
	if isExist {
		response.ToResponse("账号已存在")
		return
	}

	// 验证码生成
	err = svc.GenerateCheckCode(param.Email)
	if err != nil {
		global.Logger.ErrorF("svc.GenerateCheckCode err: %v", err)
		response.ToErrorResponse(errorCode.ErrorGenerateCheckCodeFail)
		return
	}

	// 进行响应
	response.ToResponse("验证码发送成功")
}
