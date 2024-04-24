package svc

import (
	"accessToken_grpc/internal/config"
	"accessToken_grpc/internal/db"
	"accessToken_grpc/internal/domain"
	"accessToken_grpc/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Mapper domain.TokenDao
	Token  domain.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqldb, err := db.InitMysqlDB(&c.DB.MysqlDB)
	if err != nil {
		panic("初始化数据库失败")
	}
	mapper := model.NewTokenDetailFactory(mysqldb)
	return &ServiceContext{
		Config: c,
		Mapper: mapper,
		Token:  *domain.NewTokenFactory(c.Jwt.Key, c.Jwt.Issuer, c.Jwt.Token.Exp),
	}
}
