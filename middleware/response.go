package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

var (
	SuccessCode       ResponseCode = 200
	ParamerErr        ResponseCode = 10001
	InternalErrorCode ResponseCode = 500
	NoLoginCode       ResponseCode = 20001
)

// 返回信息格式
type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
}

// 返回前端错误信息
func ResponseError(c *gin.Context, code ResponseCode, err error) {
	var (
		resp    Response
		respone []byte
	)

	resp = Response{
		ErrorCode: code,
		ErrorMsg:  err.Error(),
		Data:      "",
		TraceId:   "",
	}
	c.JSON(200, resp)
	respone, _ = json.Marshal(resp)
	c.Set("response", string(respone))
	_ = c.AbortWithError(200, err)
	c.Abort()
}

// 返回前端信息
func ResponseSuccess(c *gin.Context, data interface{}) {
	var (
		response []byte
		resp     Response
	)

	resp = Response{
		ErrorCode: SuccessCode,
		ErrorMsg:  "",
		Data:      data,
		TraceId:   "",
	}
	c.JSON(200, resp)
	response, _ = json.Marshal(resp)
	c.Set("response", string(response))
	c.Next()
}
