package articlesAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/internal/service/validator"
	"Goose/pkg/app"
	"Goose/pkg/convert"
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

//Get
// @Summary  	获取单个标签
// @Description	获取单个标签
// @Tags	 	标签管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	name       query     string         true  "标签名称"  maxlength(100)
// @Param    	state      query     int           false  "状态"    Enums(0, 1)  default(1)
// @Success  	200        {object}  model.Tag      "成功"
// @Failure  	400        {object}  errorCode.Error  "请求错误"
// @Failure  	500        {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/tags [get]
func (tag Tag) Get(context *gin.Context) {
	// 参数校验
	param := validator.GetTagRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取标签
	svc := service.New(context.Request.Context())
	newTag, err := svc.GetTag(param.Name, param.State)
	if err != nil {
		global.Logger.ErrorF("svc.GetTag err: %v", err)
		response.ToErrorResponse(errorCode.ErrorGetTagFail)
		return
	}

	response.ToResponse(newTag)
	return
}

//List
// @Summary  	获取多个标签
// @Description	获取多个标签
// @Tags	 	标签管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	name       query     string         false  "标签名称"  maxlength(100)
// @Param    	state      query     int            false  "状态"    Enums(0, 1)  default(1)
// @Param    	page       query     int            false  "页码"
// @Param    	page_size  query     int            false  "每页数量"
// @Success  	200        {object}  model.Tag      "成功"
// @Failure  	400        {object}  errorCode.Error  "请求错误"
// @Failure  	500        {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/tags/multi [get]
func (tag Tag) List(context *gin.Context) {
	// 参数校验
	param := validator.TagListRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取页码
	svc := service.New(context.Request.Context())
	totalRows, err := svc.CountTag(param.Name, param.State)
	if err != nil {
		global.Logger.ErrorF("svc.CountTag err: %v", err)
		response.ToErrorResponse(errorCode.ErrorCountTagFail)
		return
	}

	// 获取列表
	pager := app.Pager{Page: app.GetPage(context), PageSize: app.GetPageSize(context)}
	tags, err := svc.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
	if err != nil {
		global.Logger.ErrorF("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errorCode.ErrorGetTagListFail)
		return
	}

	// 进行响应
	response.ToResponseList(tags, totalRows)
	return
}

//Create
// @Summary  	新增标签
// @Description	创建一个新标签
// @Tags	 	标签管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	name        query      string         true   "标签名称"  minlength(1)  maxlength(100)
// @Param    	state       query      int            false  "状态"    Enums(0, 1)   default(1)
// @Param    	created_by  query      string         true   "创建者"   minlength(1)  maxlength(100)
// @Success  	200         {object}  model.Tag      "成功"
// @Failure  	400         {object}  errorCode.Error  "请求错误"
// @Failure  	500         {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/tags [post]
func (tag Tag) Create(context *gin.Context) {
	// 参数校验
	param := validator.CreateTagRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 创建标签
	svc := service.New(context.Request.Context())
	err := svc.CreateTag(param.Name, param.State, param.CreatedBy)
	if err != nil {
		global.Logger.ErrorF("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errorCode.ErrorCreateTagFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}

//Update
// @Summary  	更新标签
// @Description	更新标签
// @Tags	 	标签管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	id           query      int            true   "标签 ID"
// @Param    	name         query   	string         false  "标签名称"  minlength(1)  maxlength(100)
// @Param    	state        query      int            false  "状态"    Enums(0, 1)   default(1)
// @Param    	modified_by  formData   string         true   "修改者"   minlength(1)  maxlength(100)
// @Success  	200          {array}   model.Tag      "成功"
// @Failure  	400          {object}  errorCode.Error  "请求错误"
// @Failure  	500          {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/tags [put]
func (tag Tag) Update(context *gin.Context) {
	// 参数校验
	param := validator.UpdateTagRequest{ID: convert.StrTo(context.Param("id")).MustUInt32()}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 更新标签
	svc := service.New(context.Request.Context())
	err := svc.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
	if err != nil {
		global.Logger.ErrorF("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errorCode.ErrorUpdateTagFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}

//Delete
// @Summary  	删除标签
// @Description	删除标签
// @Tags	 	标签管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	id   query      int            true  "标签 ID"
// @Success  	200  {string}  string         "成功"
// @Failure  	400  {object}  errorCode.Error  "请求错误"
// @Failure  	500  {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/tags [delete]
func (tag Tag) Delete(context *gin.Context) {
	// 参数校验
	param := validator.DeleteTagRequest{ID: convert.StrTo(context.Param("id")).MustUInt32()}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 删除标签
	svc := service.New(context.Request.Context())
	err := svc.DeleteTag(param.ID)
	if err != nil {
		global.Logger.ErrorF("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errorCode.ErrorDeleteTagFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}
