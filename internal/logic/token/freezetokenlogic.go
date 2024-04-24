package token

import (
	"context"

	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreezeTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFreezeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreezeTokenLogic {
	return &FreezeTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FreezeTokenLogic) FreezeToken(req *types.TokenReq) (resp *types.TokenResp, err error) {
	err = l.svcCtx.Mapper.FreezeToken(l.ctx, req.Token)
	if err != nil {
		return nil, err
	}
	resp = types.Success()
	return resp, nil
}
