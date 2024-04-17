package handler

import (
	"accessToken/global"
	"accessToken/internal/service"
	"accessToken/model/common/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type JwtHandler struct {
}

func NewJwtApi() *JwtHandler {
	return &JwtHandler{}
}

// 链路追踪 traceId tdd
func (jwt *JwtHandler) GenToken(ctx *gin.Context) {
	var jwtService = service.NewJwtService(ctx, global.GAL_DB)
	userid := ctx.Query("userid")
	id, err := strconv.Atoi(userid)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	token, err := jwtService.GenerateToken(ctx, id)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, token)
}
