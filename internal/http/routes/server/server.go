package server

import (
	"github.com/gin-gonic/gin"

	serverController "chopipay/internal/http/controllers/server"
)

// PingExample godoc
// @Summary Server health check
// @Description Check if the server is running
// @Tags server
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", serverController.CheckServerHealth)
}