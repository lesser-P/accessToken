package system

import (
	"errors"
)

type TokenDetails struct {
	BaseModel
	UserId   int    `json:"userId" gorm:"column:user_id"`
	Token    string `json:"token" gorm:"column:token"`
	IsFreeze bool   `json:"isFreeze" gorm:"column:is_freeze"`
}

func (t *TokenDetails) TableName() string {
	return "token_details"
}

func (t *TokenDetails) SaveToDB(td *TokenDetails) error {
	result := td.DB.Save(td)
	if result.Error != nil {
		return errors.New("保存Token详情到数据库失败,err:" + result.Error.Error())
	}
	return nil
}

func (t *TokenDetails) UpdateModel(td *TokenDetails) error {
	result := t.DB.Updates(td)
	if result.Error != nil {
		return errors.New("修改Token详情到数据库失败,err:" + result.Error.Error())
	}
	return nil
}

func (t *TokenDetails) DeleteById(id string) error {
	return nil
}
