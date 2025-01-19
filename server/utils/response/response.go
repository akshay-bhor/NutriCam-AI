package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"data":    data,
		"message": message,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error":   true,
		"data":    data,
		"message": message,
	})
}
