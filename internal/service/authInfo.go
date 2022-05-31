package service

import (
	"Goose/global"
	"Goose/internal/model"
	"Goose/pkg/errorCode"
	"mime/multipart"
)

//ModifyAuthInfo 修改用户信息
func (svc *Service) ModifyAuthInfo(nickName, desc, authName string) *errorCode.Error {
	err := svc.dao.ModifyAuthInfo(authName, nickName, desc)
	if err != nil {
		global.Logger.ErrorF("svc.ModifyAuthInfo err: %v", err)
		return errorCode.ErrorModifyAuthInfo.WithDetails(err.Error())
	}
	return nil
}

//ModifyAvatar 修改用户头像
func (svc *Service) ModifyAvatar(authName string, file multipart.File, fileHeader *multipart.FileHeader) *errorCode.Error {
	// Upload Avatar
	fileInfo, err := svc.UploadAvatar(file, fileHeader)
	if err != nil {
		return errorCode.ErrorUploadAvatar.WithDetails(err.Error())
	}

	// Modify database
	err = svc.dao.ModifyAvatar(authName, fileInfo.AccessUrl)
	if err != nil {
		global.Logger.ErrorF("svc.ModifyAuthInfo err: %v", err)
		return errorCode.ErrorModifyAuthInfo.WithDetails(err.Error())
	}
	return nil
}

//GetAuthInfo 获取用户信息
func (svc *Service) GetAuthInfo(authName string) (*model.AuthInfo, *errorCode.Error) {
	// Modify database
	auth, err := svc.dao.GetAuthByAuthName(authName)
	if err != nil {
		global.Logger.ErrorF("svc.GetAuthByAuthName err: %v", err)
		return nil, errorCode.ErrorGetAuthInfo.WithDetails(err.Error())
	}
	return &model.AuthInfo{
		AuthName:     auth.AuthName,
		NickName:     auth.NickName,
		Desc:         auth.Desc,
		HeadImageUrl: auth.HeadImageUrl,
	}, nil
}
