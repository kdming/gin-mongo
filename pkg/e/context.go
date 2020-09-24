package e

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func NewError(msg string) error {
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
