package config

import (
	"accessToken_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		types.MysqlDB
	}
	Jwt struct {
		types.Token
	}
}
