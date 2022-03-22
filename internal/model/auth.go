package model

type Auth struct {
	*Model
	AuthName string `json:"auth_name"`
	AuthCode string `json:"auth_code"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func (a Auth) TableName() string {
	return "goose_auth"
}
