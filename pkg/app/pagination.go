package app

//分页处理

import (
	"Goose/global"
	"Goose/pkg/convert"
	"github.com/gin-gonic/gin"
)

//GetPage 获取分页
func GetPage(context *gin.Context) int {
	page := convert.StrTo(context.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

//GetPageSize 获取分页大小
func GetPageSize(context *gin.Context) int {
	pageSize := convert.StrTo(context.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

//GetPageOffset 获取分页偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
