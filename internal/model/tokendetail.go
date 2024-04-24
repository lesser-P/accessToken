package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type TokenDetail struct {
	BaseModel
	Token     string    `json:"token" gorm:"column:token"`
	UserId    string    `json:"userId" gorm:"column:user_id"`
	IsFreeze  bool      `json:"isFreeze" gorm:"column:is_freeze"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"column:expires_at"`
}

func (t *TokenDetail) TableName() string {
	return "token_details"
}

func NewTokenDetail(db *gorm.DB) *TokenDetail {
	return &TokenDetail{
		BaseModel: BaseModel{DB: db},
	}
}
func NewBaseTokenDetail(token, userId string, freeze bool, exp time.Time) *TokenDetail {
	return &TokenDetail{
		Token:     token,
		UserId:    userId,
		IsFreeze:  freeze,
		ExpiresAt: exp,
		BaseModel: BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: 0,
		},
	}
}

func NewTokenDetailFactory(ctx context.Context, db *gorm.DB) *TokenDetail {
	return &TokenDetail{
		BaseModel: BaseModel{Ctx: ctx, DB: db},
	}
}

func (t *TokenDetail) SaveOrUpdate(ctx context.Context, info *TokenDetail) error {
	if info.ID == nil {
		if tx := t.DB.Table("token_details").Create(info); tx.Error != nil {
			return tx.Error
		}
	} else {
		if tx := t.DB.Updates(info).Model(info); tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (t *TokenDetail) DeleteByUserId(ctx context.Context, userId string) error {
	if err := t.DB.Where("user_id=?", userId).Delete(&TokenDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *TokenDetail) SelectByToken(ctx context.Context, token string) (*TokenDetail, error) {
	result := new(TokenDetail)
	if tx := t.DB.Where("token=?", token).Select(result); tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}

func (t *TokenDetail) SelectByUserId(ctx context.Context, userId string) (*TokenDetail, error) {
	result := new(TokenDetail)
	if tx := t.DB.Where("user_id=?", userId).Select(result); tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}

func (t *TokenDetail) FreezeToken(ctx context.Context, token string) error {
	result := new(TokenDetail)
	t.DB.Where("token = ?", token).Take(result)
	if result.IsFreeze {
		return errors.New("该Token已被冻结")
	}

	result.IsFreeze = true
	if tx := t.DB.Save(result); tx.Error != nil {
		return tx.Error
	}
	return nil
}
