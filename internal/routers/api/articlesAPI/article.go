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

type Article struct{}

func NewArticle() Article {
	return Article{}
}

//Get
// @Summary  	获取单个文章
// @Description	获取单个文章
// @Tags	 	文章管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	title      	query     	string  	false  	"文章名称"  	maxlength(100)
// @Param		created_by	query		string		false	"文章作者"	maxlength(100)
// @Param    	state      	query     	int         false  	"状态"    Enums(0, 1)  default(1)
// @Success  	200        {object}  model.Article      "成功"
// @Failure  	400        {object}  errorCode.Error  "请求错误"
// @Failure  	500        {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/article/{title} [get]
func (article Article) Get(context *gin.Context) {
	// 参数校验
	param := validator.GetArticleRequest{Title: convert.StrTo(context.Param("title")).String()}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取文章
	svc := service.New(context.Request.Context())
	newArticle, err := svc.GetArticle(param.Title, param.CreatedBy, param.State)
	if err != nil {
		global.Logger.ErrorF("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errorCode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(newArticle)
	return
}

//List
// @Summary  	获取多个文章
// @Description	获取多个文章
// @Tags	 	文章管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	title      query     string         false  "文章名称"  maxlength(100)
// @Param    	state      query     int            false  "状态"    Enums(0, 1)  default(1)
// @Param    	page       query     int            false  "页码"
// @Param    	page_size  query     int            false  "每页数量"
// @Success  	200        {object}  model.Article      "成功"
// @Failure  	400        {object}  errorCode.Error  "请求错误"
// @Failure  	500        {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/article [get]
func (article Article) List(context *gin.Context) {
	// 参数校验
	param := validator.ArticleListRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 获取页码
	svc := service.New(context.Request.Context())
	totalRows, err := svc.CountArticle(param.Title, param.State)
	if err != nil {
		global.Logger.ErrorF("svc.CountArticle err: %v", err)
		response.ToErrorResponse(errorCode.ErrorCountArticleFail)
		return
	}

	// 获取列表
	pager := app.Pager{Page: app.GetPage(context), PageSize: app.GetPageSize(context)}
	articles, err := svc.GetArticleList(param.Title, param.State, pager.Page, pager.PageSize)
	if err != nil {
		global.Logger.ErrorF("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errorCode.ErrorGetArticleListFail)
		return
	}

	// 进行响应
	response.ToResponseList(articles, totalRows)
	return
}

//Create
// @Summary  	创建文章
// @Description	创建文章
// @Tags	 	文章管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	title       	formData    string	true   	"文章标题"  	minlength(1)  maxlength(100)
// @Param		desc			formData	string	true	"文章描述"	minlength(1)  maxlength(250）
// @Param		content			formData	string	true	"文章内容"	minlength(1)
// @Param		cover_image_url	formData	string	false	"封面路径"	maxlength(100)
// @Param    	state       	formData    int       false  	"状态"    	Enums(0, 1)   default(1)
// @Param    	created_by  	formData    string    true   	"创建者"   	minlength(1)  maxlength(100)
// @Success  	200  {string}  string         "成功"
// @Failure  	400  {object}  errorCode.Error  "请求错误"
// @Failure  	500  {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/article [POST]
func (article Article) Create(context *gin.Context) {
	// 参数校验
	param := validator.CreateArticleRequest{}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 创建文章
	svc := service.New(context.Request.Context())
	err := svc.CreateArticle(
		param.Title,
		param.Desc,
		param.Content,
		param.CoverImageUrl,
		param.CreatedBy,
		param.State,
	)
	if err != nil {
		global.Logger.ErrorF("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errorCode.ErrorCreateArticleFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}

//Update
// @Summary  	更新文章
// @Description	更新文章
// @Tags	 	文章管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	id           	path      	int     true   	"文章 ID"
// @Param    	title       	formData    string	false   "文章标题"  	maxlength(100)
// @Param		desc			formData	string	false	"文章描述"	maxlength(250）
// @Param		content			formData	string	false	"文章内容"
// @Param		cover_image_url	formData	string	false	"封面路径"	maxlength(100)
// @Param    	state        	formData    int     false  "状态"    Enums(0, 1)   default(1)
// @Param    	modified_by  	formData    string  true   "修改者"   minlength(3)  maxlength(100)
// @Success  	200          {array}   model.Article      "成功"
// @Failure  	400          {object}  errorCode.Error  "请求错误"
// @Failure  	500          {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/article/{id} [put]
func (article Article) Update(context *gin.Context) {
	// 参数校验
	param := validator.UpdateArticleRequest{
		ID: convert.StrTo(context.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 更新标签
	svc := service.New(context.Request.Context())
	err := svc.UpdateArticle(
		param.ID,
		param.Title,
		param.Desc,
		param.Content,
		param.CoverImageUrl,
		param.ModifiedBy,
		param.State,
	)
	if err != nil {
		global.Logger.ErrorF("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errorCode.ErrorUpdateArticleFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}

//Delete
// @Summary  	删除文章
// @Description	删除文章
// @Tags	 	文章管理
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	id   path      int            true  "文章 ID"
// @Success  	200  {string}  string         "成功"
// @Failure  	400  {object}  errorCode.Error  "请求错误"
// @Failure  	500  {object}  errorCode.Error  "内部错误"
// @Router   	/api/v1/article/{id} [delete]
func (article Article) Delete(context *gin.Context) {
	// 参数校验
	param := validator.DeleteArticleRequest{ID: convert.StrTo(context.Param("id")).MustUInt32()}
	response := app.NewResponse(context)
	valid, errs := app.BindAndValid(context, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 删除标签
	svc := service.New(context.Request.Context())
	err := svc.DeleteArticle(param.ID)
	if err != nil {
		global.Logger.ErrorF("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errorCode.ErrorDeleteArticleFail)
		return
	}

	// 进行响应
	response.ToResponse(gin.H{})
	return
}
