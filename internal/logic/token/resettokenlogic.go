package token

import (
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
}

func NewResetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetTokenLogic {
	return &ResetTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetTokenLogic) ResetToken(req *types.TokenReq) (resp *types.TokenResp, err error) {
	resp = &types.TokenResp{}
	token := req.Token
	data, err := l.svcCtx.Mapper.SelectByToken(l.ctx, token)
	if err != nil {
		return nil, err
	}

	if !l.svcCtx.Token.VailToken(l.ctx, req.Token) {
		return nil, errors.New("验证不通过无法重置")
	}

	newToken, err := l.svcCtx.Token.ResetToken(l.ctx, token)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	data.Token = newToken
	data.ExpiresAt = time.Now().Add(time.Hour * time.Duration(l.svcCtx.Config.Jwt.Timeout))
	err = l.svcCtx.Mapper.SaveOrUpdate(l.ctx, data)
	if err != nil {
		return nil, err
	}
	resp.Data = newToken
	resp.Code = 200
	return resp, nil
}
