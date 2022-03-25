package dao

import (
	"Goose/internal/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	db := d.engine

	// auth校验
	db = db.Where("binary auth_name = ? AND binary auth_code = ? AND is_del = ?", appKey, appSecret, 0)

	// auth获取
	var auth model.Auth
	db = db.First(&auth)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

func (d *Dao) GetAuthByEmail(email string) (model.Auth, error) {
	db := d.engine

	// auth校验
	db = db.Where("binary email = ? AND is_del = ?", email, 0)

	// auth获取
	var auth model.Auth
	db = db.First(&auth)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

//GenerateCheckCode 生成验证码
func (pool *Pool) GenerateCheckCode(key string, val string) error {
	conn := pool.Pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("Set", key, val)
	_, _ = conn.Do("expire", key, 300) // 五分钟
	if err != nil {
		return err
	}

	return nil
}
