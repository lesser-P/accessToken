package main

import (
	"accessToken/global"
	"accessToken/initialized"
)

// TODO context贯穿全文
// TODO 理解引用透明性原则
// TODO 实现go-mock单元测试

func main() {
	global.GAL_Viper = initialized.NewViper()
	initialized.NewMysqlDB()
	initialized.RunServer()
}
