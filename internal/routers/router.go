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
		apiAuth.POST("/login", auth.Login)
		apiAuth.POST("/sendCheck", auth.SendCheck)
		apiAuth.POST("/register", auth.Register)
		apiAuth.PUT("/modifyCode", auth.ModifyCode)
		apiAuth.PATCH("/modifyCode", auth.ModifyCode)
		apiAuth.PUT("/resetCode", auth.ResetCode)
		apiAuth.PATCH("/resetCode", auth.ResetCode)
	}

	// 注册服务API
	article := articlesAPI.NewArticle()
	tag := articlesAPI.NewTag()
	authInfo := authAPI.NewAuthInfo()
	apiService := router.Group("/api/v1")
	apiService.Use(middleware.JWT())
	{
		apiService.PUT("/auth/info", authInfo.Modify)
		apiService.PATCH("/auth/info", authInfo.Modify)
		apiService.PUT("/auth/avatar", authInfo.ModifyAvatar)
		apiService.PATCH("/auth/avatar", authInfo.ModifyAvatar)
		apiService.GET("/auth/info", authInfo.Get)

		apiService.POST("/tags", tag.Create)
		apiService.DELETE("/tags", tag.Delete)
		apiService.PUT("/tags", tag.Update)
		apiService.PATCH("/tags/state", tag.Update)
		apiService.GET("/tags", tag.Get)
		apiService.GET("/tags/multi", tag.List)

		apiService.POST("/article", article.Create)
		apiService.DELETE("/article", article.Delete)
		apiService.PUT("/article", article.Update)
		apiService.PATCH("/article/state", article.Update)
		apiService.GET("/article", article.Get)
		apiService.GET("/article/multi", article.List)
	}

	return router
}
