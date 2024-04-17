package routers

import (
	"accessToken/internal/handler"
	"github.com/gin-gonic/gin"
)

type JwtRouterGroupInit struct {
}

func (jwt *JwtRouterGroupInit) InitRouterWithAuth(group *gin.RouterGroup) {
	group.Group("jwt")
	{
		group.GET("token", handler.JwtHandler{}.GenToken)
	}
}
func (jwt *JwtRouterGroupInit) InitRouterWithNoAuth(group *gin.RouterGroup) {

}
