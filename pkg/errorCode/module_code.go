package errorCode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
	ErrorGetTagFail     = NewError(20010006, "获取标签失败")

	ErrorGetArticleListFail = NewError(20020001, "获取文章列表失败")
	ErrorCreateArticleFail  = NewError(20020002, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(20020003, "更新文章失败")
	ErrorDeleteArticleFail  = NewError(20020004, "删除文章失败")
	ErrorCountArticleFail   = NewError(20020005, "统计文章失败")
	ErrorGetArticleFail     = NewError(20020006, "获取文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")

	ErrorGenerateCheckCodeFail = NewError(20040002, "生成验证码失败")
	ErrorFormatEmail           = NewError(20040003, "邮箱格式错误")

	ErrorFormatAuthName    = NewError(20040005, "账号格式错误")
	ErrorFormatAuthCode    = NewError(20040006, "密码格式错误")
	ErrorNotValidCheckCode = NewError(20040007, "验证码对应错误")
	ErrorCreateNewAuth     = NewError(20040008, "账号创建失败")
	ErrorAuthNameExist     = NewError(20040009, "账号名已存在")
	ErrorEmailExist        = NewError(20040010, "邮箱已存在")

	ErrorModifyCode = NewError(20040011, "修改密码失败")
)
