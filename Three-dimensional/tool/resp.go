package tool

import "github.com/gin-gonic/gin"

func ReturnFailure(ctx *gin.Context, info interface{}) {
	ctx.JSON(200, gin.H{
		"msg":  200,
		"info": info,
	})
}

func ReturnSuccess(ctx *gin.Context, info interface{}) {
	ctx.JSON(200, gin.H{
		"msg":  200,
		"info": info,
	})
}
