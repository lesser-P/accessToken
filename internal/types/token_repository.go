package types

import (
	"accessToken_go_zero/internal/model"
	"context"
)

type TokenDao interface {
	SaveOrUpdate(ctx context.Context, info *model.TokenDetail) error
	DeleteByUserId(ctx context.Context, userId string) error
	SelectByToken(ctx context.Context, token string) (*model.TokenDetail, error)
	SelectByUserId(ctx context.Context, userId string) (*model.TokenDetail, error)
	FreezeToken(ctx context.Context, token string) error
}
