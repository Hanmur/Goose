package routers

import (
	_ "Goose/docs"
	"Goose/global"
	"Goose/internal/middleware"
	"Goose/internal/routers/api/articlesAPI"
	"Goose/internal/routers/api/authAPI"
	"Goose/internal/routers/api/uploadAPI"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)


//NewRouter 路由设置
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.Translations())

	// 注册Swagger的获取API
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册静态资源访问的API
	router.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 注册文件上传的API
	apiUpload := router.Group("/upload")
	apiUpload.Use(middleware.JWT()).POST("/file", uploadAPI.NewUpload().UploadFile)

	// 注册Token认证的API
	router.POST("/auth", authAPI.GetAuth)

	// 注册文章管理API
	article := articlesAPI.NewArticle()
	tag := articlesAPI.NewTag()
	apiArticle := router.Group("/api/articles")
	apiArticle.Use(middleware.JWT())
	{
		apiArticle.POST("/tags", tag.Create)
		apiArticle.DELETE("/tags/:id", tag.Delete)
		apiArticle.PUT("/tags/:id", tag.Update)
		apiArticle.PATCH("/tags/:id/state", tag.Update)
		apiArticle.GET("/tags/:id", tag.Get)
		apiArticle.GET("/tags", tag.List)

		apiArticle.POST("/article", article.Create)
		apiArticle.DELETE("/article/:id", article.Delete)
		apiArticle.PUT("/article/:id", article.Update)
		apiArticle.PATCH("/article/:id/state", article.Update)
		apiArticle.GET("/article/:id", article.Get)
		apiArticle.GET("/article", article.List)
	}

	return router
}

