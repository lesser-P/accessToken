package initialized

import (
	"accessToken/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() *gorm.DB {
	dbInfo := global.GAL_Config.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       dbInfo.DNS(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic("open DB err:" + err.Error())
	}
	return db
}
