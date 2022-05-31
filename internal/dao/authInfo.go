package dao

import "Goose/internal/model"

//ModifyAuthInfo 修改账户信息
func (d *Dao) ModifyAuthInfo(authName, nickName, desc string) error {
	auth := model.Auth{
		AuthName: authName,
		Model:    &model.Model{},
	}
	values := map[string]interface{}{
		"nick_name":   nickName,
		"desc":        desc,
		"modified_by": authName,
	}

	db := d.engine.Model(auth)
	// 查找账户
	db = db.Where("binary auth_name = ? AND is_del = ?", auth.AuthName, 0)
	// 更新账户
	db = db.Updates(values)

	return db.Error
}

//ModifyAvatar 修改账户头像
func (d *Dao) ModifyAvatar(authName, headImageUrl string) error {
	auth := model.Auth{
		AuthName: authName,
		Model:    &model.Model{},
	}
	values := map[string]interface{}{
		"head_image_url": headImageUrl,
		"modified_by":    authName,
	}

	db := d.engine.Model(auth)
	// 查找账户
	db = db.Where("binary auth_name = ? AND is_del = ?", auth.AuthName, 0)
	// 更新账户
	db = db.Updates(values)

	return db.Error
}
