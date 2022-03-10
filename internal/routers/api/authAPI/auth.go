package authAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
)


//GetAuth
// @Summary  	获取Token
// @Description	获取Token（临时使用的API）
// @Tags	 	Token验证
// @Produce  	json
// @Param    	app_key       	formData     string   	true  	"认证账号" 	Enum("Hanmur")
// @Param    	app_secret  	formData     string   	true  	"认证密码" 	Enum("Hanmur_goose)
// @Success  	200        {object}  string      		"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/auth [POST]
func GetAuth(c *gin.Context) {
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

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.ErrorF("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errorCode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
