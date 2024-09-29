package server

import (
	"github.com/gin-gonic/gin"

	serverController "chopipay/internal/http/controllers/server"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", serverController.CheckServerHealth)
}