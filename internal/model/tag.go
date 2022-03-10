package model

//Tag 标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

//TableName 获取标签数据表名
func (tag Tag) TableName() string {
	return "goose_tags"
}

