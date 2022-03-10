package dao

import (
	"Goose/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	db := d.engine

	// auth校验
	db = db.Where("app_key = ? AND app_secret = ? AND is_del = ?", appKey, appSecret, 0)

	// auth获取
	var auth model.Auth
	db = db.First(&auth)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}