package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register)
	engine.POST("/login", login)

	accountGroup := engine.Group("/account")
	{
		accountGroup.Use(auth)
		accountGroup.POST("/createAccount",createAccount)
		accountGroup.POST("/recharge",recharge)
	}

	paymentGroup:= engine.Group("/transferMoney")
	{
		paymentGroup.Use(auth)
		paymentGroup.POST("/transferMoney",transferMoney)
		paymentGroup.GET("/selectDetails",selectDetails)
	}

	engine.Run()
}
