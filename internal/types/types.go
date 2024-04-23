// Code generated by goctl. DO NOT EDIT.
package types

type BaseResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Code    int    `json:"code"`
}

type GenTokenReq struct {
	UserId string `json:"userId"`
}

type TokenReq struct {
	Token  string `form:"token",json:"token"`
	UserId string `form:"userId",json:"userId"`
}

type TokenResp struct {
	BaseResponse
}