package service

import (
	"accessToken/global"
	"accessToken/internal/dao"
	"accessToken/pkg/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JwtService struct {
	dao dao.TokenDao
}

func NewJwtService(ctx *gin.Context, db *gorm.DB) *JwtService {
	return &JwtService{
		dao: dao.NewJwtDao(db),
	}
}

var jwtInfo = global.GAL_Config.Jwt

const (
	EMPTY = ""
)

// 生成Token
func (jwt *JwtService) GenerateToken(ctx *gin.Context, userId int) (string, error) {
	record, err := jwt.dao.GetRecordById(ctx, userId)
	if err != nil {
		return EMPTY, err
	}
	token, createErr := auth.CreateToken(ctx, userId, false, jwtInfo.Issuer, jwtInfo.Key)
	if createErr != nil {
		return EMPTY, createErr
	}
	if record != nil {
		tk, tokenErr := auth.ParseReturnToken(ctx, record.Token, jwtInfo.Key)
		if tokenErr != nil {
			return EMPTY, tokenErr
		}
		if !tk.Valid {
			err = jwt.dao.CreateTokenRecord(ctx, token, userId)
			if err != nil {
				return EMPTY, err
			}
			return token, nil
		}
		return EMPTY, errors.New("该用户已有token")
	}
	return token, nil
}

// 判断token是否冻结
func (jwt *JwtService) IsFreeze() {

}
