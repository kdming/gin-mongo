package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Cors 跨域处理
func Cors(api *gin.RouterGroup) {
	api.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "OPTIONS", "POST", "DELETE"},
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: false,
		Debug:              true,
	}))
	api.OPTIONS("*options_support", func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNoContent)
		return
	})
}
