package logic

import (
	"context"

	"accessToken_grpc/internal/svc"
	"accessToken_grpc/pb/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenLogic) VerifyToken(in *token.TokenReq) (*token.TokenResp, error) {
	err := l.svcCtx.Token.VailToken(l.ctx, in.Token)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	return token.Success(), nil
}
