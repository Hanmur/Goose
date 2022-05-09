package errorCode

import "net/http"

//内部错误码定义表
var (
	Success          = NewError(00000000, http.StatusOK, "成功")
	ServerError      = NewError(10000000, http.StatusInternalServerError, "服务内部错误")
	InvalidParams    = NewError(10000001, http.StatusBadRequest, "入参错误")
	NotFound         = NewError(10000002, http.StatusNotFound, "找不到")
	TooManyRequests  = NewError(10000003, http.StatusTooManyRequests, "请求过多")
	ParamsTokenError = NewError(10000004, http.StatusBadRequest, "入参Token缺失")
)
