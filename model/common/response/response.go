package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS = 200
	ERROR   = 500
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSONP(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, nil, "操作成功", c)
}

func OkWithMessage(c *gin.Context, msg string) {
	Result(SUCCESS, nil, msg, c)
}

func OkWithData(c *gin.Context, data any) {
	Result(SUCCESS, data, "操作成功", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, nil, "操作失败", c)
}

func FailWithMessage(c *gin.Context, msg string) {
	Result(ERROR, nil, msg, c)
}

func FailWithData(c *gin.Context, data any) {
	Result(ERROR, data, "操作失败", c)
}
