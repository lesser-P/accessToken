package logic

import (
	"context"
	"errors"

	"accessToken_grpc/internal/svc"
	"accessToken_grpc/pb/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreezeTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFreezeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FreezeTokenLogic {
	return &FreezeTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FreezeTokenLogic) FreezeToken(in *token.TokenReq) (*token.TokenResp, error) {
	data, err := l.svcCtx.Mapper.SelectByToken(l.ctx, in.Token)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	if data.IsFreeze {
		return token.Failed(), errors.New("该Token已被冻结")
	}
	data.IsFreeze = true
	err = l.svcCtx.Mapper.SaveOrUpdate(l.ctx, data)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	return token.Success(), nil
}
