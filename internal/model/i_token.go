package model

type TokenDao interface {
	SaveOrUpdate(info *TokenDetail) error
	DeleteByUserId(userId string) error
	SelectByToken(token string) (*TokenDetail, error)
	SelectByUserId(userId string) (*TokenDetail, error)
	FreezeToken(token string) error
}
