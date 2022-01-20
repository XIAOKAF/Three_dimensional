package api

import (
	"Three_dimensional/model"
	"Three_dimensional/service"
	"Three_dimensional/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	time2 "time"
)

func transferMoney(context *gin.Context) {
	payer, err := context.Cookie("name")
	if err != nil {
		tool.ReturnFailure(context, "转账失败")
		fmt.Println(err)
		return
	}
	payee := context.PostForm("payee")
	amount := context.PostForm("amount")
	postscript := context.PostForm("postscript")

	a, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		tool.ReturnFailure(context, "转账失败")
		fmt.Println(err)
		return
	}

	currentTime := time2.Now()

	transfer := model.Transfer{
		Payer:      payer,
		Payee:      payee,
		Amount:     float32(a),
		Postscript: postscript,
		Time:       currentTime,
	}

	err,flag := service.TransferMoney(transfer)
	if err != nil {
		tool.ReturnFailure(context,"转账失败")
		fmt.Println(err)
		return
	}
	if flag != true {
		tool.ReturnFailure(context,"余额不足")
		return
	}
	tool.ReturnFailure(context,"转账成功")
}

func selectDetails(context *gin.Context)  {
	details := context.PostForm("details")
	fmt.Println(details)
	transfer := model.Transfer{
		Postscript: details,
	}
	t,err:=service.SelectDetails(transfer)
	if err != nil {
		tool.ReturnFailure(context,"查询失败")
		fmt.Println(err)
		return
	}
	tool.ReturnSuccess(context,t)
}
