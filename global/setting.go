package global

import (
	"Goose/pkg/logger"
	"Goose/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting *setting.JWTSettingS
	Logger 	*logger.Logger
)
