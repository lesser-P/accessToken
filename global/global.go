package global

import (
	"accessToken/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GAL_Config     configs.Config
	GAL_Router     *gin.Engine
	GAL_MysqlDB    *gorm.DB
	GAL_Viper      *viper.Viper
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

const (
	RAND_START = 10000000000
	RAND_END   = 999999999999
)
