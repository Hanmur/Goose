package errorCode

import "net/http"

var (
	ErrorUploadFileFail = NewError(20030001, http.StatusInternalServerError, "上传文件失败")
)
