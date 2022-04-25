package dao

import (
	"Goose/internal/model"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

//GetAuth 账户密码校验
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

//GetAuthByAuthName 获取账户
func (d *Dao) GetAuthByAuthName(authName string) (model.Auth, error) {
	db := d.engine

	// auth校验
	db = db.Where("binary auth_name = ? AND is_del = ?", authName, 0)

	// auth获取
	var auth model.Auth
	db = db.First(&auth)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}

	return auth, nil
}

//GetAuthByEmail 通过邮箱获取账户
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
func (pool *Pool) GenerateCheckCode(key, val string) error {
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

//CheckCheckCode 验证码校验
func (pool *Pool) CheckCheckCode(key, val string) error {
	conn := pool.Pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	trueCode, err := redis.String(conn.Do("Get", key))
	if err != nil {
		return err
	}

	if trueCode != val {
		return errors.New("验证码对应错误")
	}

	return nil
}

//CreateAuth 创建新账户
func (d *Dao) CreateAuth(authName, authCode, email string) error {
	auth := model.Auth{
		AuthName: authName,
		AuthCode: authCode,
		Email:    email,
		Model:    &model.Model{CreatedBy: authName},
	}
	db := d.engine

	// 创建标签
	db = db.Create(&auth)

	return db.Error
}

//ModifyCode 修改账户密码
func (d *Dao) ModifyCode(authName, newCode string) error {
	auth := model.Auth{
		AuthName: authName,
		Model:    &model.Model{},
	}
	values := map[string]interface{}{
		"auth_code":   newCode,
		"modified_by": authName,
	}

	db := d.engine.Model(auth)
	// 查找账户
	db = db.Where("binary auth_name = ? AND is_del = ?", auth.AuthName, 0)
	// 更新账户
	db = db.Updates(values)

	return db.Error
}
