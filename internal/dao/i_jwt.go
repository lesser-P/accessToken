package dao

import (
	"accessToken/model/system"
	"gorm.io/gorm"
)

type TokenDao interface {
	CreateTokenRecord(token string, userId int) error
	GetRecordById(userId int) (*system.TokenDetails, error)
}

func NewJwtDao(db *gorm.DB) TokenDao {
	return &JwtDao{
		db: db,
	}
}
