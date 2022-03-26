package authAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/internal/service/validator"
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
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

	// 发送验证码
	svc := service.New(c.Request.Context())
	err := svc.SendCheck(param.Email)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	// 进行响应
	response.ToResponse("验证码发送成功")
}

//Register
// @Summary  	注册账号
// @Description	检验验证码和账号密码格式，进行登录
// @Tags	 	账户管理
// @Produce  	json
// @Param    	auth_name   	formData     string   	true  	"账号"
// @Param    	auth_code   	formData     string   	true  	"密码"
// @Param    	email   		formData     string   	true  	"邮箱"
// @Param    	check_code   	formData     string   	true  	"验证码"
// @Success  	200        {object}  string      		"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/auth/register [POST]
func (auth Auth) Register(c *gin.Context) {
	// 参数校验
	param := validator.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 进行注册
	svc := service.New(c.Request.Context())
	err := svc.Register(param.AuthName, param.AuthCode, param.Email, param.CheckCode)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	// 进行响应
	response.ToResponse("注册成功")
}
