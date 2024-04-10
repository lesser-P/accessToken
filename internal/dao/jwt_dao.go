package dao

import (
	"accessToken/model/system"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

//var db = global.GAL_DB

const (
	EMPTY = ""
)

type JwtDao struct {
	db *gorm.DB
}

func (jwt *JwtDao) CreateTokenRecord(token string, userId int) error {
	result := jwt.db.Create(&system.TokenDetails{
		Token:  token,
		UserId: userId,
	})
	if result.Error != nil {
		return errors.New("保存信息异常：err" + result.Error.Error())
	}
	return nil
}

func (jwt *JwtDao) GetRecordById(userId int) (*system.TokenDetails, error) {
	detail := system.TokenDetails{}
	jwt.db.Where("user_id = ?", strconv.Itoa(userId)).Take(&detail)
	if detail.Token == EMPTY {
		return nil, errors.New("数据不存在")
	}
	return &detail, nil
}
