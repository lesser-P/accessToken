package dao

import (
	"accessToken/model/system"
	"gorm.io/gorm"
)

type Dao interface {
	CreateTokenRecord(token string, userId int) error
	GetRecordById(userId int) (*system.TokenDetails, error)
}

func NewDao(db *gorm.DB) Dao {
	return &JwtDao{
		db: db,
	}
}
