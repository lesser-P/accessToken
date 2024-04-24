package logic

import (
	"accessToken_grpc/internal/model"
	"context"
	"time"

	"accessToken_grpc/internal/svc"
	"accessToken_grpc/pb/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenTokenLogic) GenToken(in *token.GenReq) (*token.TokenResp, error) {
	jwt, err := l.svcCtx.Token.GenerateToken(l.ctx, in.UserId, false)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	data := model.NewBaseToken(jwt, in.UserId, false,
		time.Now().Add(time.Hour*time.Duration(l.svcCtx.Config.Jwt.Token.Exp)))
	err = l.svcCtx.Mapper.SaveOrUpdate(l.ctx, data)
	if err != nil {
		return token.FailedWithMsgAndData(err.Error(), nil), err
	}
	return token.SuccessWithData(jwt), nil
}
