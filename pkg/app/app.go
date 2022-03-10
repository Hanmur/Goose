package app

//响应处理

import (
	"Goose/pkg/errorCode"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Response 响应包
type Response struct {
	Context *gin.Context
}

//Pager 分页器
type Pager struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

//NewResponse 创建新响应
func NewResponse(context *gin.Context) *Response {
	return &Response{Context: context}
}

//ToResponse 将数据读入为成功响应包
func (response *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	response.Context.JSON(http.StatusOK, data)
}

//ToResponseList 将数据读入为分页响应包
func (response *Response) ToResponseList(list interface{}, totalRows int) {
	tempResponse := gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(response.Context),
			PageSize:  GetPageSize(response.Context),
			TotalRows: totalRows,
		},
	}

	response.Context.JSON(http.StatusOK, tempResponse)
}

//ToErrorResponse 将错误信息读入为错误响应包
func (response *Response) ToErrorResponse(err *errorCode.Error) {
	tempResponse := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		tempResponse["details"] = details
	}

	response.Context.JSON(err.StatusCode(), tempResponse)
}