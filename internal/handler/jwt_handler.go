package handler

import (
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

func (jwt *JwtHandler) GenToken(c *gin.Context) {
	var jwtService = service.NewJwtService()
	userid := c.Query("userid")
	id, err := strconv.Atoi(userid)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	token, err := jwtService.GenerateToken(id)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithData(c, token)
}
