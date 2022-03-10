package dao

import (
	"Goose/internal/model"
	"Goose/pkg/app"
)

//CountTag 获取标签数
func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	db := d.engine.Model(&tag)

	// 检索名称为name的标签
	if name != "" {
		db = db.Where("name = ?", name)
	}
	// 检索状态为state的标签
	db = db.Where("state = ?", state)
	// 检索是删除状态为0(未删除的)的标签
	db = db.Where("is_del = ?", 0)

	// 计数
	var count int
	db.Count(&count)

	// 错误检定
	if db.Error != nil {
		return 0, db.Error
	}

	return count, nil
}

//GetTag 获取标签
func (d *Dao) GetTag(name string, state uint8) (*model.Tag, error) {
	tag := model.Tag{Name:name, State: state}
	db := d.engine.Model(&tag)

	// 检索名称为name的标签
	if tag.Name != "" {
		db = db.Where("name = ?", name)
	}
	// 检索状态为state的标签
	db = db.Where("state = ?", state)
	// 检索删除状态为0(未删除)的标签
	db = db.Where("is_del = ?", 0)
	// 录入标签
	db = db.Find(&tag)

	// 错误检定
	if db.Error != nil{
		return nil, db.Error
	}

	return &tag, nil
}

//GetTagList 获取标签列表
func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	db := d.engine.Model(&tag)

	// 设置所在页码
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// 检索名称为name的标签
	if tag.Name != "" {
		db = db.Where("name = ?", tag.Name)
	}
	// 检索状态为state的标签
	db = db.Where("state = ?", tag.State)
	// 检索删除状态为0(未删除)的标签
	db = db.Where("is_del = ?", 0)
	// 录入数据
	var tags []*model.Tag
	db = db.Find(&tags)

	// 错误检定
	if db.Error != nil {
		return nil, db.Error
	}

	return tags, nil
}

//CreateTag 创建新标签
func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	db := d.engine

	// 创建标签
	db = db.Create(&tag)

	return db.Error
}

//UpdateTag 更新标签
func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	db := d.engine.Model(tag)

	// 查找标签
	db = db.Where("id = ? AND is_del = ?", tag.ID, 0)
	// 更新标签
	db = db.Updates(values)

	return db.Error
}

//DeleteTag 删除标签
func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	db := d.engine.Model(tag)

	// 检索标签
	db = db.Where("id = ? AND is_del = ?", tag.Model.ID, 0)
	// 删除标签
	db = db.Delete(&tag)

	return db.Error
}
