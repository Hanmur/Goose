package service

import (
	"Goose/global"
	"Goose/pkg/upload"
	"Goose/pkg/util"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}

	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}

func (svc *Service) UploadAvatar(file multipart.File, fileHeader *multipart.FileHeader, authName string) (*FileInfo, error) {
	// FileName
	fileExt := upload.GetFileExt(fileHeader.Filename)
	fileName := util.EncodeMD5(authName) + fileExt

	if !upload.CheckContainExt(1, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckMaxSize(1, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	// FilePath
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	dst := uploadSavePath + "/" + "avatars" + "/" + fileName

	// FileSave
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
