package service

import (
	"Goose/internal/model"
)


func (svc *Service) CountTag(name string, state uint8) (int, error) {
	return svc.dao.CountTag(name, state)
}

func (svc *Service) GetTag(name string, state uint8) (*model.Tag, error){
	return svc.dao.GetTag(name, state)
}

func (svc *Service) GetTagList(name string, state uint8, page int, pageSize int) ([]*model.Tag, error) {
	return svc.dao.GetTagList(name, state, page, pageSize)
}

func (svc *Service) CreateTag(name string, state uint8, createdBy string) error {
	return svc.dao.CreateTag(name, state, createdBy)
}

func (svc *Service) UpdateTag(ID uint32, name string, state uint8, modifiedBy string) error {
	return svc.dao.UpdateTag(ID, name, state, modifiedBy)
}

func (svc *Service) DeleteTag(ID uint32) error {
	return svc.dao.DeleteTag(ID)
}
