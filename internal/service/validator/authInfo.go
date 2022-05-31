package validator

type ModifyInfoRequest struct {
	NickName string `form:"nick_name" binding:"max=20"`
	Desc     string `form:"desc" binding:"max=50"`
}
