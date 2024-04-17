package dao

import (
	"accessToken/model/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenDao interface {
	CreateTokenRecord(ctx *gin.Context, token string, userId int) error
	GetRecordById(ctx *gin.Context, userId int) (*system.TokenDetails, error)
}

func NewJwtDao(db *gorm.DB) TokenDao {
	return &JwtDao{
		db: db,
	}
}
