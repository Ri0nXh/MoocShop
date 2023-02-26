package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// RespSuccess 成功
func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// RespSuccess 成功
func RespSuccessWithoutData(c *gin.Context) {
	c.JSON(http.StatusOK, CommonResponse{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: nil,
	})
}

// RespSuccess 失败
func RespError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, CommonResponse{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// RespSuccess 自定义失败消息
func RespErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
