package main

import (
	"app/api"
	"app/dao"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	if !dao.Connect() {
		panic("数据库连接失败!!!")
	}

	g := gin.Default()
	api.InitRouter(g)
	g.Static("/static", "./static")

	err := g.Run(":1688")
	if err != nil {
		fmt.Println("服务启动失败" + err.Error())
	}

}
