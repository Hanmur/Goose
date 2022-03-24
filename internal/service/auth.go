package service

import (
	"Goose/internal/service/validator"
	"errors"
)

//CheckAuth 确认Auth是否存在
func (svc *Service) CheckAuth(param *validator.AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AuthName, param.AuthCode)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist")
}
