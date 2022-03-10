package service

import "Goose/internal/model"

func (svc *Service) CountArticle(title string, state uint8) (int, error) {
	return svc.dao.CountArticle(title, state)
}

func (svc *Service) GetArticle(title string, createdBy string, state uint8) (*model.Article, error){
	return svc.dao.GetArticle(title, createdBy, state)
}

func (svc *Service) GetArticleList(title string, state uint8, page int, pageSize int) ([]*model.Article, error) {
	return svc.dao.GetArticleList(title, state, page, pageSize)
}

func (svc *Service) CreateArticle(title, desc, content, coverImageUrl, createdBy string, state uint8) error {
	return svc.dao.CreateArticle(title, desc, content, coverImageUrl, createdBy, state)
}

func (svc *Service) UpdateArticle(ID uint32, title, desc, content, coverImageUrl, modifiedBy string, state uint8) error {
	return svc.dao.UpdateArticle(ID, title, desc, content, coverImageUrl, modifiedBy, state)
}

func (svc *Service) DeleteArticle(ID uint32) error {
	return svc.dao.DeleteArticle(ID)
}
