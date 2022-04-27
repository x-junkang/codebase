package handler

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, &Response{Code: 200, Msg: "success", Data: data})
}

func ResponseFail(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, &Response{Code: code, Msg: msg})
}
