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
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenTokenLogic) GenToken(req *types.GenTokenReq) (resp string, err error) {
	token, err := l.svcCtx.Token.GenerateToken(l.ctx, req.UserId, false)
	if err != nil {
		return "", err
	}
	data := model.NewBaseTokenDetail(token, req.UserId, false, time.Now().Add(time.Hour*time.Duration(l.svcCtx.Config.Jwt.Timeout)))
	err = l.svcCtx.Mapper.SaveOrUpdate(l.ctx, data)
	if err != nil {
		return "", err
	}
	return token, nil
}
