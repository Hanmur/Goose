package authAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
)

type Auth struct{}

func NewAuth() Auth {
	return Auth{}
}

//CheckIn
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
func (auth Auth) CheckIn(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.ErrorF("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errorCode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AuthName, param.AuthCode)
	if err != nil {
		global.Logger.ErrorF("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errorCode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
