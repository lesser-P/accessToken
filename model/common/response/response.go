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

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(ctx *gin.Context) {
	Result(SUCCESS, nil, "操作成功", ctx)
}

func OkWithMessage(ctx *gin.Context, msg string) {
	Result(SUCCESS, nil, msg, ctx)
}

func OkWithData(ctx *gin.Context, data any) {
	Result(SUCCESS, data, "操作成功", ctx)
}

func Fail(ctx *gin.Context) {
	Result(ERROR, nil, "操作失败", ctx)
}

func FailWithMessage(ctx *gin.Context, msg string) {
	Result(ERROR, nil, msg, ctx)
}

func FailWithData(ctx *gin.Context, data any) {
	Result(ERROR, data, "操作失败", ctx)
}
