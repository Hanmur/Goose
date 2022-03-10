package model

//Article 文章
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

//TableName 获取文章数据表名
func (article Article) TableName() string {
	return "goose_article"
}
