package errorshandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error, errMsg string) {
	logError := fmt.Sprintf("%s: %s", errMsg, err.Error())
	log.Println(logError)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errMsg,
	})
}
