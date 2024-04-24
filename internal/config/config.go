package config

import (
	"accessToken_go_zero/internal/model"
	"accessToken_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		model.MysqlDB
	}
	Jwt struct {
		types.Token
	}
}
