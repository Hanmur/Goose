package service

import "errors"

type AuthRequest struct {
	AuthName string `form:"auth_name" binding:"required"`
	AuthCode string `form:"auth_code" binding:"required"`
}

//CheckAuth 确认Auth是否存在
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AuthName, param.AuthCode)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
