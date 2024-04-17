package main

import (
	"accessToken/global"
	"accessToken/initialized"
	"accessToken/routers"
)

func main() {
	routers.InitApiRouter()
}

type A interface {
	Get() string
}

type Handler struct {
}

func (h *Handler) Get() string {
	return ""
}

var _ A = (*Handler)(nil)

func init() {
	global.GAL_Viper = initialized.NewViper()
	global.GAL_DB = initialized.NewMysqlDB()
	initialized.RunServer()

}
