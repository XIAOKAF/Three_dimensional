package api

import (
	"Three_dimensional/model"
	"Three_dimensional/service"
	"Three_dimensional/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func register(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("pwd")

	user := model.User{
		Name:     name,
		Password: password,
	}

	flag, err := service.IsNameRepeated(name)
	if err != nil {
		tool.ReturnFailure(context, "注册失败")
		fmt.Println(err)
		return
	}
	if flag != true {
		tool.ReturnFailure(context, "用户名已存在")
		return
	}

	userId, err := service.Register(user)
	if err != nil {
		tool.ReturnFailure(context, "注册失败")
		fmt.Println(err)
		return
	}

	id := strconv.Itoa(userId)
	tool.ReturnSuccess(context, "注册成功啦,你的账号是："+id)
}

func login(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("pwd")

	user := model.User{
		Name: name,
		Password: password,
	}

	flag, err := service.IsNameExist(name)
	if err != nil {
		tool.ReturnFailure(context, "登录失败")
		fmt.Println(err)
		return
	}
	if flag != true {
		tool.ReturnFailure(context, "账号不存在，先注册一下吧")
		return
	}

	pwd, err := service.SelectPasswordByUserName(user)
	if err != nil {
		tool.ReturnFailure(context, "登录失败")
		fmt.Println(err)
		return
	}
	if pwd != password {
		tool.ReturnFailure(context, "密码错误")
		return
	}
	context.SetCookie("name", name, 600, "/", "localhost", false, true)
	tool.ReturnSuccess(context, "登录成功")
}
