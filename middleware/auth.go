package middleware

import (
	"app/common/config"
	"app/models"
	"app/service/user_service"

	"github.com/gin-gonic/gin"
)

// Auth 身份验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.JSON(401, gin.H{"code": 1, "msg": "token不能为空"})
			c.Abort()
			return
		}

		tokenSvc := &user_service.TokenSvc{}
		user, err := tokenSvc.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{"code": 1, "msg": err.Error()})
			c.Abort()
			return
		}

		c.Set("userId", user.Id)
		c.Set("role", user.Role)

		// 记录请求日志
		if config.GetConfig().SaveRequestLog {
			logger := &models.Logger{}
			logger.Start()
			logger.UserId = user.Id
			defer logger.End(c)
		}
		c.Next()
	}
}
