package types

import (
	"accessToken_go_zero/pkg/auth"
	"context"
	"errors"
)

type Token struct {
	Key     string `json:"key"`
	Issuer  string `json:"issuer"`
	Timeout int    `json:"timeout"`
}

func CreateToken(key, issuer string, timeout int) *Token {
	return &Token{
		Key:     key,
		Issuer:  issuer,
		Timeout: timeout,
	}
}

func (t *Token) GenerateToken(ctx context.Context, userId string, isFreeze bool) (string, error) {
	token, err := auth.CreateToken(ctx, userId, isFreeze, t.Issuer, t.Key)
	if err != nil {
		return "", errors.New("生成Token失败，err:" + err.Error())
	}
	return token, nil
}

func (t *Token) VailToken(ctx context.Context, token string) bool {
	_, err := auth.ParseToken(ctx, token, t.Key)
	if err != nil {
		return false
	}
	return true
}

func (t *Token) ResetToken(ctx context.Context, token string) (string, error) {
	token, err := auth.RefreshTokenStatus(ctx, token, t.Key)
	if err != nil {
		return "", err
	}
	return token, nil
}
