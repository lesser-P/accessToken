package config

import (
	"accessToken_grpc/internal/db"
	"accessToken_grpc/internal/domain"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Jwt struct {
		domain.Token
	}
	DB struct {
		db.MysqlDB
	}
}
