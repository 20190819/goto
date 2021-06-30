package utils

import "github.com/gin-gonic/gin"

func Success() {

}

func ErrorResponse(ctx *gin.Context, status int, code int, msg string) {
	ctx.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"msg":  "请求成功",
		"data": data,
	})
}
