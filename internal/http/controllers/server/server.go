package server

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func CheckServerHealth(c *gin.Context) {
	log.Println("Checking server health...")
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}