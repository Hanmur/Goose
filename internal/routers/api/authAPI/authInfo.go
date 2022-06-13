package authAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/internal/service/validator"
	"Goose/pkg/app"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
)

type AuthInfo struct{}

func NewAuthInfo() AuthInfo {
	return AuthInfo{}
}

//Modify
// @Summary  	修改个人信息
// @Description	修改个人信息
// @Tags	 	账户管理
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param    	nick_name   	formData     string   	false  	"别名"
// @Param    	desc		   	formData     string   	false  	"描述"
// @Success  	200        {object}  string	 	     	"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/api/v1/auth/info [PUT]
func (authInfo AuthInfo) Modify(c *gin.Context) {
	// 参数校验
	param := validator.ModifyInfoRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取账户
	authName := c.MustGet("auth_name").(string)

	// 修改信息
	svc := service.New(c.Request.Context())
	err := svc.ModifyAuthInfo(param.NickName, param.Desc, authName)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	// 进行响应
	response.ToResponse("修改信息成功")
}

//ModifyAvatar
// @Summary  	修改头像
// @Description	修改头像
// @Tags	 	账户管理
// @Produce  	json
// @Security 	ApiKeyAuth
// @Param    	head_image  	formData     file   	true  	"头像"
// @Success  	200        {object}  string	 	     	"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/api/v1/auth/avatar [PUT]
func (authInfo AuthInfo) ModifyAvatar(c *gin.Context) {
	response := app.NewResponse(c)

	// 获取文件
	file, fileHeader, errF := c.Request.FormFile("head_image")
	if errF != nil {
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errF.Error()))
		return
	}

	// 获取账户
	authName := c.MustGet("auth_name").(string)

	// 修改头像
	svc := service.New(c.Request.Context())
	err := svc.ModifyAvatar(authName, file, fileHeader)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	// 进行响应
	response.ToResponse("修改头像成功")
}

//Get
// @Summary  	获取个人信息
// @Description	获取个人信息
// @Tags	 	账户管理
// @Produce  	json
// @Security 	ApiKeyAuth
// @Success  	200        {object}  string	 	     	"成功"
// @Failure  	400        {object}  errorCode.Error  	"请求错误"
// @Failure  	500        {object}  errorCode.Error  	"内部错误"
// @Router   	/api/v1/auth/info [GET]
func (authInfo AuthInfo) Get(c *gin.Context) {
	response := app.NewResponse(c)
	// 获取账户
	authName := c.MustGet("auth_name").(string)
	global.Logger.Info(authName)
	// 修改信息
	svc := service.New(c.Request.Context())
	newAuthInfo, err := svc.GetAuthInfo(authName)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}
	global.Logger.Info(newAuthInfo.AuthName)
	// 进行响应
	response.ToResponse(newAuthInfo)
	return
}
