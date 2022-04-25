package validator

type LoginRequest struct {
	AuthName string `form:"auth_name" binding:"required,min=4,max=20"`
	AuthCode string `form:"auth_code" binding:"required,min=4,max=20"`
}

type SendCheckRequest struct {
	Email  string `form:"email" binding:"required,min=1,max=50"`
	Prefix string `form:"prefix" binding:"oneof=regis reset"`
}

type RegisterRequest struct {
	AuthName  string `form:"auth_name" binding:"required,min=4,max=20"`
	AuthCode  string `form:"auth_code" binding:"required,min=4,max=20"`
	Email     string `form:"email" binding:"required,min=1,max=50"`
	CheckCode string `form:"check_code" binding:"required,min=6,max=6"`
}

type ModifyCodeRequest struct {
	AuthName string `form:"auth_name" binding:"required,min=4,max=20"`
	AuthCode string `form:"auth_code" binding:"required,min=4,max=20"`
	NewCode  string `form:"new_code" binding:"required,min=4,max=20"`
}

type ResetCodeRequest struct {
	Email     string `form:"email" binding:"required,min=1,max=50"`
	CheckCode string `form:"check_code" binding:"required,min=6,max=6"`
}
