package api

import (
	"Three_dimensional/model"
	"Three_dimensional/service"
	"Three_dimensional/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func createAccount(context *gin.Context) {
	accountName, err := context.Cookie("name")
	if err != nil {
		tool.ReturnFailure(context, "账户创建失败")
		fmt.Println(err)
		return
	}
	err = service.CreateAccount(accountName)
	if err != nil {
		tool.ReturnFailure(context, "账户创建失败")
		fmt.Println(err)
		return
	}
	tool.ReturnSuccess(context, "成功创建账户")
}

func recharge(context *gin.Context) {
	accountName, err := context.Cookie("name")
	if err != nil {
		tool.ReturnFailure(context, "充值失败")
		fmt.Println(err)
		return
	}
	income := context.PostForm("recharge")

	i, err := strconv.ParseFloat(income, 32)
	if err != nil {
		tool.ReturnFailure(context, "充值失败")
		fmt.Println(err)
		return
	}

	account := model.Account{
		AccountName: accountName,
		Income:      float32(i),
	}

	err, total := service.Recharge(account)
	if err != nil {
		tool.ReturnFailure(context, "充值失败")
		fmt.Println(err)
		return
	}
	sum := strconv.FormatFloat(float64(total),'E',-1,32)
	tool.ReturnSuccess(context,"成功充值"+income+"元，"+"目前你的余额是:"+sum)
}
