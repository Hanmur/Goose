package errorCode

import "net/http"

var (
	ErrorGetTagListFail = NewError(20010001, http.StatusInternalServerError, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, http.StatusInternalServerError, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, http.StatusInternalServerError, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, http.StatusInternalServerError, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, http.StatusInternalServerError, "统计标签失败")
	ErrorGetTagFail     = NewError(20010006, http.StatusInternalServerError, "获取标签失败")
	ErrorTagExist       = NewError(20010007, http.StatusBadRequest, "标签已存在")
)
