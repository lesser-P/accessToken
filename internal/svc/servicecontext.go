package svc

import (
	"accessToken_go_zero/internal/config"
	"accessToken_go_zero/internal/types"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB, err := types.InitMysqlDB(&c.DB.MysqlDB)
	if err != nil {
		panic("初始化失败")
	}
	return &ServiceContext{
		Config: c,
		DB:     mysqlDB,
	}
}
