package token

import (
	"net/http"
)

func Success() *TokenResp {
	return &TokenResp{
		Code:    http.StatusOK,
		Message: "成功",
	}
}
func SuccessWithData(data string) *TokenResp {
	return &TokenResp{
		Code:    http.StatusOK,
		Message: "成功",
		Data:    data,
	}
}
func SuccessWithMsgAndData(msg string, data string) *TokenResp {
	return &TokenResp{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	}
}

func Failed() *TokenResp {
	return &TokenResp{
		Code:    http.StatusBadRequest,
		Message: "失败",
	}
}

func FailedWithData(data any) *TokenResp {
	return &TokenResp{
		Code:    http.StatusBadRequest,
		Message: "失败",
	}
}

func FailedWithMsgAndData(msg string, data any) *TokenResp {
	return &TokenResp{
		Code:    http.StatusBadRequest,
		Message: "失败：" + msg,
	}
}
