package model

type Auth struct {
	*Model
	AuthName     string `json:"auth_name"`
	AuthCode     string `json:"auth_code"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Desc         string `json:"desc"`
	HeadImageUrl string `json:"head_image_url"`
}

type AuthInfo struct {
	AuthName     string `json:"auth_name"`
	NickName     string `json:"nick_name"`
	Desc         string `json:"desc"`
	HeadImageUrl string `json:"head_image_url"`
}

func (a Auth) TableName() string {
	return "goose_auth"
}
