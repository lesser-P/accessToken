package token

import (
	"accessToken_go_zero/internal/model"
	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	mapper model.TokenDao
}

func NewResetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext, mapper model.TokenDao) *ResetTokenLogic {
	return &ResetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		mapper: mapper,
	}
}

func (l *ResetTokenLogic) ResetToken(req *types.TokenReq) (resp *types.TokenResp, err error) {
	resp = &types.TokenResp{}
	token := req.Token
	data, err := l.mapper.SelectByToken(token)
	if err != nil {
		return nil, err
	}
	jwt := types.CreateToken(l.svcCtx.Config.Jwt.Key, l.svcCtx.Config.Jwt.Issuer, l.svcCtx.Config.Jwt.Timeout)
	if !jwt.VailToken(l.ctx, req.Token) {
		return nil, errors.New("验证不通过无法重置")
	}

	newToken, err := jwt.ResetToken(l.ctx, token)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	data.Token = newToken
	data.ExpiresAt = time.Now().Add(time.Hour * time.Duration(l.svcCtx.Config.Jwt.Timeout))
	err = l.mapper.SaveOrUpdate(data)
	if err != nil {
		return nil, err
	}
	resp.Data = newToken
	resp.Code = 200
	return resp, nil
}
