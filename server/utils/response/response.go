package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorObject struct {
	code     int
	response gin.H
}

func SuccessResponse(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"data":    data,
		"message": message,
	})
}

func NewErrorResponse(statusCode int, data interface{}, message string) ErrorObject {
	return ErrorObject{
		code: statusCode,
		response: gin.H{
			"error":   true,
			"data":    data,
			"message": message,
		},
	}
}

func ErrorResponse(c *gin.Context, errObj ErrorObject) {
	c.AbortWithStatusJSON(errObj.code, errObj.response)
}
