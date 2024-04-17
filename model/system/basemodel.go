package system

import "gorm.io/gorm"

type BaseModel struct {
	*gorm.DB
	ID         int    `json:"id" gorm:"column:id"`
	CreateTime string `json:"createTime" gorm:"column:create_time"`
	UpdateTime string `json:"updateTime" gorm:"column:update_time"`
	IsDeleted  int    `json:"isDeleted" gorm:"column:is_deleted"`
}
