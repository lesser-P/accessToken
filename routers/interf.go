package routers

import "github.com/gin-gonic/gin"

type IRouterGroupInit interface {
	InitRouterWithAuth(group *gin.RouterGroup)
	InitRouterWithNoAuth(group *gin.RouterGroup)
}
