package errorCode

import "net/http"

var (
	ErrorGetArticleListFail = NewError(20020001, http.StatusInternalServerError, "获取文章列表失败")
	ErrorCreateArticleFail  = NewError(20020002, http.StatusInternalServerError, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(20020003, http.StatusInternalServerError, "更新文章失败")
	ErrorDeleteArticleFail  = NewError(20020004, http.StatusInternalServerError, "删除文章失败")
	ErrorCountArticleFail   = NewError(20020005, http.StatusInternalServerError, "统计文章失败")
	ErrorGetArticleFail     = NewError(20020006, http.StatusInternalServerError, "获取文章失败")
	ErrorUpdateArticlePower = NewError(20020007, http.StatusUnauthorized, "更新文章无权限")
	ErrorDeleteArticlePower = NewError(20020008, http.StatusUnauthorized, "删除文章无权限")
	ErrorArticleNotFound    = NewError(20020009, http.StatusBadRequest, "文章不存在")
)
