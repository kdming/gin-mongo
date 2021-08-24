package middleware

import (
	"github.com/gin-gonic/gin"
)

// 异常捕获
func CustomError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if c.IsAborted() {
				c.Status(200)
			}

			switch errStr := err.(type) {
			case string:
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": errStr})
			case error:
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": "系统错误" + errStr.Error()})
			default:
				c.JSON(200, gin.H{"code": 1, "data": nil, "msg": "发生未知异常"})
			}
		}
	}()
	c.Next()
}
