package token

import (
	"context"
	"errors"

	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(req *types.TokenReq) (resp *types.TokenResp, err error) {
	jwt := types.CreateToken(l.svcCtx.Config.Jwt.Key, l.svcCtx.Config.Jwt.Issuer, l.svcCtx.Config.Jwt.Timeout)
	if !jwt.VailToken(l.ctx, req.Token) {
		return nil, errors.New("验证失败")
	}
	resp = &types.TokenResp{
		BaseResponse: types.BaseResponse{
			Message: "验证通过",
			Code:    200,
		},
	}
	return resp, nil
}
