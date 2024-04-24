package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	Path     string `json:"path" yaml:"path"`
	Port     string `json:"port" yaml:"port"`
	UserName string `json:"username" yaml:"username"`
	Pwd      string `json:"password" yaml:"password"`
	DbName   string `json:"dbname" yaml:"dbname"`
}

func InitMysqlDB(m *MysqlDB) (*gorm.DB, error) {
	config := mysql.Config{
		DSN:                       m.DSN(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: true,
	}
	db, err := gorm.Open(mysql.New(config))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (m *MysqlDB) DSN() string {
	s := m.UserName + ":" + m.Pwd + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	return s
}
