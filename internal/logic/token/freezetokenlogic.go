package token

import (
	"accessToken_go_zero/internal/model"
	"context"

	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FreezeTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	mapper model.TokenDao
}

func NewFreezeTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext, mapper model.TokenDao) *FreezeTokenLogic {
	return &FreezeTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		mapper: mapper,
	}
}

func (l *FreezeTokenLogic) FreezeToken(req *types.TokenReq) (resp *types.TokenResp, err error) {
	err = l.mapper.FreezeToken(req.Token)
	if err != nil {
		return nil, err
	}
	resp = &types.TokenResp{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "冻结成功",
		},
	}
	return resp, nil
}
