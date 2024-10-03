package errorshandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err error, errMsg string) {
	errMsg = fmt.Sprintf("%s: %s", errMsg, err.Error())
	log.Println(errMsg)
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": errMsg,
	})
}