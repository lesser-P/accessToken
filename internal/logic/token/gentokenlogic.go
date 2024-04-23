package token

import (
	"accessToken_go_zero/internal/model"
	"context"
	"time"

	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	mapper model.TokenDao
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext, mapper model.TokenDao) *GenTokenLogic {
	return &GenTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		mapper: mapper,
	}
}

func (l *GenTokenLogic) GenToken(req *types.GenTokenReq) (resp string, err error) {
	tk := types.CreateToken(l.svcCtx.Config.Jwt.Key, l.svcCtx.Config.Jwt.Issuer, l.svcCtx.Config.Jwt.Timeout)
	token, err := tk.GenerateToken(l.ctx, req.UserId, false)
	if err != nil {
		return "", err
	}
	data := model.TokenDetail{
		UserId:    req.UserId,
		Token:     token,
		IsFreeze:  false,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(l.svcCtx.Config.Jwt.Timeout)),
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			IsDeleted: 0,
		},
	}
	err = l.mapper.SaveOrUpdate(&data)
	if err != nil {
		return "", err
	}
	return token, nil
}
