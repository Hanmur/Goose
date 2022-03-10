package uploadAPI

import (
	"Goose/global"
	"Goose/internal/service"
	"Goose/pkg/app"
	"Goose/pkg/convert"
	"Goose/pkg/errorCode"
	"Goose/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

//UploadFile
// @Summary  	文件上传
// @Description	文件上传，目前仅支持图片
// @Tags	 	上传
// @Produce  	json
// @Security ApiKeyAuth
// @Param    	file       	formData  	file         		true  	"文件路径"
// @Param		type		formData	int					true	"文件类型"	Enums(1)
// @Success  	200        	{object}	string      		"成功"
// @Failure  	400        	{object}  	errorCode.Error  	"请求错误"
// @Failure  	500        	{object}  	errorCode.Error  	"内部错误"
// @Router   	/upload/file [POST]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errorCode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errorCode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.ErrorF("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errorCode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
