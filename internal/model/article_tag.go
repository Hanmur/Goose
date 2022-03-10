package model

//ArticleTag 文章标签关联器
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

//TableName 获取文章标签关联数据表名
func (a ArticleTag) TableName() string {
	return "goose_tag_article"
}
