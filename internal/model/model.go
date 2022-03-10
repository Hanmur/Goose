package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Model 通用文本创建模型
type Model struct{
	ID			uint32 `gorm:"primary_key" json:"id"`
	CreatedBy	string `json:"created_by"`
	ModifiedBy	string `json:"modified_by"`
	CreatedOn  	uint32 `json:"created_on"`
	ModifiedOn 	uint32 `json:"modified_on"`
	DeletedOn  	uint32 `json:"deleted_on"`
	IsDel      	uint8  `json:"is_del"`
}

