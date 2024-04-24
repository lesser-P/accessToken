package svc

import (
	"accessToken_go_zero/internal/config"
	"accessToken_go_zero/internal/model"
	"accessToken_go_zero/internal/types"
	"context"
)

type ServiceContext struct {
	Config config.Config
	Mapper types.TokenDao
	Token  *types.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB, err := model.InitMysqlDB(&c.DB.MysqlDB)
	ctx := context.Background()
	if err != nil {
		panic("初始化失败")
	}
	return &ServiceContext{
		Config: c,
		Mapper: model.NewTokenDetailFactory(ctx, mysqlDB),
		Token:  types.NewTokenFactory(c.Jwt.Key, c.Jwt.Issuer, c.Jwt.Timeout),
	}
}
