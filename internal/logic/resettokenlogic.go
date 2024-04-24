package logic

import (
	"context"

	"accessToken_grpc/internal/svc"
	"accessToken_grpc/pb/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetTokenLogic {
	return &ResetTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetTokenLogic) ResetToken(in *token.TokenReq) (*token.TokenResp, error) {
	newToken, err := l.svcCtx.Token.ResetToken(l.ctx, in.Token)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	data, err := l.svcCtx.Mapper.SelectByToken(l.ctx, in.Token)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	data.Token = newToken
	err = l.svcCtx.Mapper.SaveOrUpdate(l.ctx, data)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	return token.SuccessWithData(newToken), nil
}
