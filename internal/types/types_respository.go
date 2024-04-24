package types

import "net/http"

func Success() *TokenResp {
	return &TokenResp{
		BaseResponse: *NewBaseResp("成功", http.StatusOK, nil),
	}
}
func SuccessWithData(data any) *TokenResp {
	return &TokenResp{
		BaseResponse: *NewBaseResp("成功", http.StatusOK, data),
	}
}

func Failed() *TokenResp {
	return &TokenResp{
		BaseResponse: *NewBaseResp("失败", http.StatusBadRequest, nil),
	}
}
func FailedWithMsg(msg string) *TokenResp {
	return &TokenResp{
		BaseResponse: *NewBaseResp("失败:"+msg, http.StatusBadRequest, nil),
	}
}

func NewBaseResp(msg string, code int, data any) *BaseResponse {
	return &BaseResponse{
		Message: msg,
		Code:    code,
		Data:    data,
	}
}
