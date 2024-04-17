package system

import "gorm.io/gorm"

type BaseModel struct {
	*gorm.DB
	ID         int    `json:"id"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
