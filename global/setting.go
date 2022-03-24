package global

import (
	"Goose/pkg/logger"
	"Goose/pkg/setting"
)

var (
	ServerSetting    *setting.ServerSettingS
	AppSetting       *setting.AppSettingS
	DatabaseSetting  *setting.DatabaseSettingS
	JWTSetting       *setting.JWTSettingS
	EmailSetting     *setting.EmailSettingS
	RedisPoolSetting *setting.RedisPoolSettingS
	Logger           *logger.Logger
)
