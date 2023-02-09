package v1

import (
	"app/common/app"

	"github.com/gin-gonic/gin"
)

func AuthTest(c *gin.Context) {
	app.Ok(c, "通过测试", nil)
}
