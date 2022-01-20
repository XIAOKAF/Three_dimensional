package main

import (
	"Three_dimensional/api"
	"Three_dimensional/dao"
)

func main()  {
	dao.InitDB()
	api.InitEngine()
}
