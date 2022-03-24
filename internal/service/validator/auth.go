package validator

type AuthRequest struct {
	AuthName string `form:"auth_name" binding:"required,min=4,max=20"`
	AuthCode string `form:"auth_code" binding:"required,min=4,max=20"`
}
