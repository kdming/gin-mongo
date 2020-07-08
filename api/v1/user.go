package v1

import (
	"app/models"
	"app/pkg/e"
	"app/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		e.Err("bind", err)
	}

	if err := user.Save(); err != nil {
		e.Err("注册失败", err)
	}

	tkModel := &jwt.TokenModel{}
	tkModel.User = user.Id
	token, err := jwt.MakeToken(tkModel)
	if err != nil {
		e.Err("token生成失败", err)
	}

	e.Ok(c, "注册成功", e.Map{"token": token})

}

func Login(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		e.Err("bind", err)
	}

	if err := user.Login(); err != nil {
		e.Err("登录失败", err)
	}

	tkModel := &jwt.TokenModel{}
	tkModel.User = user.Id
	token, err := jwt.MakeToken(tkModel)
	if err != nil {
		e.Err("token生成失败", err)
	}

	e.Ok(c, "登录成功", e.Map{"token": token})

}
