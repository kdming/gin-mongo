package main

import (
	"app/api"
	"app/common/config"
	"app/dao"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	dao.Connect()

	g := gin.Default()
	api.InitRouter(g)
	g.Static("/static", "./static")

	appPort := "1688"
	if config.GetConfig().APPPORT != "" {
		appPort = config.GetConfig().APPPORT
	}
	err := g.Run(fmt.Sprintf(":%v", appPort))
	if err != nil {
		fmt.Println("服务启动失败" + err.Error())
	}
}
