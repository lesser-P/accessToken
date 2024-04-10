package global

import (
	"accessToken/configs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GAL_Config configs.Config
	GAL_Router *gin.Engine
	GAL_DB     *gorm.DB
	GAL_Viper  *viper.Viper
)
