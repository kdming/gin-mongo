package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewError(params ...interface{}) error {
	msg := ""
	for i := 0; i < len(params); i++ {
		val := params[i]
		switch val.(type) {
		case string:
			msg += val.(string)
		case error:
			if msg != "" {
				msg += ":"
			}
			msg += val.(error).Error()
		default:
			msg += "发生错误"
		}
	}
	return errors.New(msg)
}

type Map map[string]interface{}

func Ok(c *gin.Context, msg string, data interface{}) {
	res := make(map[string]interface{})
	res["code"] = 0
	res["msg"] = msg
	res["data"] = data
	c.JSON(200, res)
}

func Err(msg string, err error) {
	if err != nil {
		msg += ":" + err.Error()
	}
	panic(msg)
}

func GetUserId(c *gin.Context) primitive.ObjectID {
	userId, _ := c.Get("userId")
	return userId.(primitive.ObjectID)
}

func GetUserRole(c *gin.Context) int {
	role, _ := c.Get("role")
	return role.(int)
}
