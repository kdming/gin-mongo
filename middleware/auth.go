package middleware

import (
	"app/models"
	"app/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// 身份验证
func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("token")

		if token == "" {
			c.JSON(401, gin.H{"code": 1, "msg": "token不能为空"})
			c.Abort()
			return
		}

		tokenModel, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{"code": 1, "msg": err.Error()})
			c.Abort()
			return
		}

		c.Set("user", tokenModel.User)

		// 记录请求日志
		logger := &models.Logger{}
		logger.Start()
		logger.User = tokenModel.User

		c.Next()

		logger.End(c)

	}

}
