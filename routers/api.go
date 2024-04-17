package routers

import "github.com/gin-gonic/gin"

func InitApiRouter() {
	var router *gin.Engine
	router = gin.Default()
	initRouter(router)
}
