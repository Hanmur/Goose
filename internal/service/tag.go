package service

import (
	"Goose/global"
	"Goose/internal/model"
	"Goose/pkg/errorCode"
)

//CountTag 计算tag个数
func (svc *Service) CountTag(name string, state uint8) (int, error) {
	return svc.dao.CountTag(name, state)
}

//GetTag 获取单个标签
func (svc *Service) GetTag(name string, state uint8) (*model.Tag, error) {
	return svc.dao.GetTag(name, state)
}

//GetTagList 获取标签页
func (svc *Service) GetTagList(name string, state uint8, page int, pageSize int) ([]*model.Tag, error) {
	return svc.dao.GetTagList(name, state, page, pageSize)
}

//CreateTag 创建新标签
func (svc *Service) CreateTag(name string, state uint8, createdBy string) *errorCode.Error {
	// 重复检验
	_, err := svc.GetTag(name, state)
	if err == nil {
		return errorCode.ErrorTagExist
	} else if err.Error() != "record not found" {
		global.Logger.ErrorF("svc.GetTag err: %v", err)
		return errorCode.ErrorGetTagFail
	}

	// 创建标签
	err = svc.dao.CreateTag(name, state, createdBy)
	if err != nil {
		global.Logger.ErrorF("svc.CreateTag err: %v", err)
		return errorCode.ErrorCreateTagFail
	}

	return nil
}

//UpdateTag 更新标签
func (svc *Service) UpdateTag(ID uint32, name string, state uint8, modifiedBy string) error {
	return svc.dao.UpdateTag(ID, name, state, modifiedBy)
}

//DeleteTag 删除标签
func (svc *Service) DeleteTag(ID uint32) error {
	return svc.dao.DeleteTag(ID)
}
