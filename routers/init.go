package routers

import (
	"github.com/gin-gonic/gin"
)

func initRouter(router *gin.Engine) {
	initCommonGroup(router.Group("/common"))
}

func initCommonGroup(group *gin.RouterGroup) {
	interfs := []IRouterGroupInit{
		&JwtRouterGroupInit{},
	}
	// 不需验证的路由
	for _, interf := range interfs {
		interf.InitRouterWithNoAuth(group)
	}
	// 添加中间件
	group.Use()
	// 需要token的路由
	for _, interf := range interfs {
		interf.InitRouterWithAuth(group)
	}
}
