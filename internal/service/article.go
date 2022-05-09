package service

import (
	"Goose/global"
	"Goose/internal/model"
	"Goose/pkg/errorCode"
)

func (svc *Service) CountArticle(title string, state uint8) (int, error) {
	return svc.dao.CountArticle(title, state)
}

func (svc *Service) GetArticle(title string, createdBy string, state uint8) (*model.Article, error) {
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

func (svc *Service) DeleteArticle(ID uint32, authName string) *errorCode.Error {
	// 检验删除权限
	article, err := svc.dao.GetArticleByID(ID)
	if err != nil {
		global.Logger.ErrorF("svc.dao.GetArticleByID err: %v", err)
		return errorCode.ErrorArticleNotFound
	}
	if article.CreatedBy != authName {
		return errorCode.ErrorDeleteArticlePower
	}

	// 删除文章
	err = svc.dao.DeleteArticle(ID)
	if err != nil {
		global.Logger.ErrorF("svc.DeleteArticle err: %v", err)
		return errorCode.ErrorDeleteArticleFail
	}
	return nil
}
