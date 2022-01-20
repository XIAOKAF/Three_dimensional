package api

import (
	"Three_dimensional/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func auth(context *gin.Context) {
	name, err := context.Cookie("name")
	if err != nil {
		tool.ReturnFailure(context, "还未登录")
		fmt.Println(err)
		context.Abort()
	}
	context.Set("name", name)
	context.Next()
}
