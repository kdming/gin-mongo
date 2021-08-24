package v1

import (
	"app/common/app"
	"app/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		app.Err("bind", err)
	}

	if err := userSvc.Register(user); err != nil {
		app.Err("注册失败", err)
	}

	token, err := tokenSvc.MakeToken(user)
	if err != nil {
		app.Err("token生成失败", err)
	}
	app.Ok(c, "注册成功", app.Map{"token": token})
}

func Login(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		app.Err("bind", err)
	}

	if err := userSvc.Login(user); err != nil {
		app.Err("登录失败", err)
	}

	token, err := tokenSvc.MakeToken(user)
	if err != nil {
		app.Err("token生成失败", err)
	}
	app.Ok(c, "登录成功", app.Map{"token": token})
}
