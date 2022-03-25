package validator

type LoginRequest struct {
	AuthName string `form:"auth_name" binding:"required,min=4,max=20"`
	AuthCode string `form:"auth_code" binding:"required,min=4,max=20"`
}

type GetEmailRequest struct {
	Email string `form:"email" binding:"required,min=1,max=50"`
}
