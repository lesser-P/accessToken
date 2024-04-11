package service

import (
	"accessToken/global"
	"accessToken/internal/dao"
	"accessToken/pkg/auth"
	"errors"
)

type JwtService struct {
	dao dao.TokenDao
}

func NewJwtService() *JwtService {
	return &JwtService{
		dao: dao.NewJwtDao(global.GAL_DB),
	}
}

var jwtInfo = global.GAL_Config.Jwt

const (
	EMPTY = ""
)

// 生成Token
func (jwt *JwtService) GenerateToken(userId int) (string, error) {
	token, err := auth.CreateToken(userId, false, jwtInfo.Issuer, jwtInfo.Key)
	if err != nil {
		return EMPTY, err
	}

	record, err := jwt.dao.GetRecordById(userId)
	if err != nil {
		return EMPTY, err
	}
	if record != nil {
		return EMPTY, errors.New("该用户ID已存在token")
	}
	err = jwt.dao.CreateTokenRecord(token, userId)
	if err != nil {
		return EMPTY, err
	}
	return token, nil
}

// 判断token是否冻结
func (jwt *JwtService) IsFreeze() {

}
