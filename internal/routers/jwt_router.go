package routers

import (
	"accessToken/internal/handler"
	"github.com/gin-gonic/gin"
)

type JwtRouter struct {
}

func NewJwtRouter() *JwtRouter {
	return &JwtRouter{}
}

func (router *JwtRouter) InitJwtRouter(group *gin.RouterGroup) {
	jwtRouter := group.Group("jwt")
	jwtApi := handler.NewJwtApi()
	{
		jwtRouter.GET("/token", jwtApi.GenToken)
	}
}
