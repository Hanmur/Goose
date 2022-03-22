package routers

import (
	_ "Goose/docs"
	"Goose/global"
	"Goose/internal/middleware"
	"Goose/internal/routers/api/articlesAPI"
	"Goose/internal/routers/api/authAPI"
	"Goose/internal/routers/api/uploadAPI"
	"Goose/pkg/limiter"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
)

//methodLimiters 流量限制器
var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimitBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

//NewRouter 路由设置
func NewRouter() *gin.Engine {
	router := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		router.Use(gin.Logger())
		router.Use(gin.Recovery())
	} else {
		router.Use(middleware.AccessLog())
		router.Use(middleware.Recovery())
	}

	router.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	router.Use(middleware.RateLimiter(methodLimiters))
	router.Use(middleware.Translations())

	// 注册Swagger的获取API
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 注册静态资源访问的API
	router.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	// 注册文件上传的API
	apiUpload := router.Group("/upload")
	apiUpload.Use(middleware.JWT()).POST("/file", uploadAPI.NewUpload().UploadFile)

	// 注册Token认证的API
	auth := authAPI.NewAuth()
	apiAuth := router.Group("/auth")
	{
		apiAuth.POST("/login", auth.CheckIn)
	}

	// 注册服务API
	article := articlesAPI.NewArticle()
	tag := articlesAPI.NewTag()
	apiService := router.Group("/api/v1")
	apiService.Use(middleware.JWT())
	{
		apiService.POST("/tags", tag.Create)
		apiService.DELETE("/tags/:id", tag.Delete)
		apiService.PUT("/tags/:id", tag.Update)
		apiService.PATCH("/tags/:id/state", tag.Update)
		apiService.GET("/tags/:id", tag.Get)
		apiService.GET("/tags", tag.List)

		apiService.POST("/article", article.Create)
		apiService.DELETE("/article/:id", article.Delete)
		apiService.PUT("/article/:id", article.Update)
		apiService.PATCH("/article/:id/state", article.Update)
		apiService.GET("/article/:id", article.Get)
		apiService.GET("/article", article.List)
	}

	return router
}
