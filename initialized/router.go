package initialized

import (
	"accessToken/global"
	"accessToken/internal/middleware"
	"accessToken/routers"
	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	router := gin.Default()
	// 加载跨域请求中间件
	router.Use(middleware.Cors())

	prefix := global.GAL_Config.RouterPrefix
	group := router.Group(prefix)

	jwtRouter := routers.NewJwtRouter()
	jwtRouter.InitJwtRouter(group)

	return router
}
