package dao

import (
	"Goose/internal/model"
	"Goose/pkg/app"
)


//CountArticle 获取文章数
func (d *Dao) CountArticle(title string, state uint8) (int, error) {
	article := model.Article{Title: title, State: state}
	db := d.engine.Model(&article)

	// 检索名称为name的标签
	if title != "" {
		db = db.Where("title = ?", title)
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

//GetArticle 获取文章
func (d *Dao) GetArticle(title string, createdBy string, state uint8) (*model.Article, error) {
	tempModel := &model.Model{CreatedBy: createdBy}
	article := model.Article{Title:title, Model:tempModel, State: state}
	db := d.engine.Model(&article)

	// 检索名称为name的标签
	if title != "" {
		db = db.Where("title = ?", title)
	}
	if createdBy != ""{
		db = db.Where("created_by = ?", createdBy)
	}
	// 检索状态为state的标签
	db = db.Where("state = ?", state)
	// 检索删除状态为0(未删除)的标签
	db = db.Where("is_del = ?", 0)
	// 录入标签
	db = db.Find(&article)

	// 错误检定
	if db.Error != nil{
		return nil, db.Error
	}

	return &article, nil
}

//GetArticleList 获取文章列表
func (d *Dao) GetArticleList(title string, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{Title: title, State: state}
	db := d.engine.Model(&article)

	// 设置所在页码
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	// 检索名称为name的标签
	if article.Title != "" {
		db = db.Where("title = ?", article.Title)
	}
	// 检索状态为state的标签
	db = db.Where("state = ?", article.State)
	// 检索删除状态为0(未删除)的标签
	db = db.Where("is_del = ?", 0)
	// 录入数据
	var articles []*model.Article
	db = db.Find(&articles)

	// 错误检定
	if db.Error != nil {
		return nil, db.Error
	}

	return articles, nil
}

//CreateArticle 创建新文章
func (d *Dao) CreateArticle(title, desc, content, coverImageUrl, createdBy string, state uint8) error {
	article := model.Article{
		Title: title,
		Desc: desc,
		Content: content,
		CoverImageUrl: coverImageUrl,
		State: state,
		Model: &model.Model{CreatedBy: createdBy},
	}
	db := d.engine

	// 创建标签
	db = db.Create(&article)

	return db.Error
}

//UpdateArticle 更新文章
func (d *Dao) UpdateArticle(ID uint32, title, desc, content, coverImageUrl, modifiedBy string, state uint8) error {
	article := model.Article{
		Model: &model.Model{ID: ID},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != ""{
		values["desc"] = desc
	}
	if content != ""{
		values["content"] = content
	}
	if coverImageUrl != ""{
		values["coverImageUrl"] = coverImageUrl
	}

	db := d.engine.Model(article)

	// 查找标签
	db = db.Where("id = ? AND is_del = ?", article.ID, 0)
	// 更新标签
	db = db.Updates(values)

	return db.Error
}

//DeleteArticle 删除文章
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	db := d.engine.Model(article)

	// 检索标签
	db = db.Where("id = ? AND is_del = ?", article.Model.ID, 0)
	// 删除标签
	db = db.Delete(&article)

	return db.Error
}
