package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Ctx       context.Context
	DB        *gorm.DB
	ID        *int      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updateAt" gorm:"column:updated_at"`
	IsDeleted int       `json:"isDeleted" gorm:"column:is_deleted"`
}
