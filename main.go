package main

import (
	"Goose/global"
	"Goose/internal/dao"
	"Goose/internal/routers"
	"Goose/pkg/logger"
	"Goose/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupRedisPool()
	if err != nil {
		log.Fatalf("init.setupRedisPool err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

//Swagger设置
// @Title Goose谷声
// @Version 1.0
// @Schemes http://localhost:8000/swagger/index.html
// @Description 简单的API描述文档
// @contact.name Hanmur
// @contact.url https://hanmur.cn/
// @contact.email wenyt8@mail2.edu.cn
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}

func setupSetting() error {
	sets, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = sets.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = sets.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = sets.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = sets.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = sets.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = sets.ReadSection("RedisPool", &global.RedisPoolSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = dao.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupRedisPool() error {
	var err error
	global.RedisPool, err = dao.NewRedisEngine(global.RedisPoolSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	var err error
	global.Logger, err = logger.NewLogger(
		global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
	)
	if err != nil {
		return err
	}

	return nil
}
