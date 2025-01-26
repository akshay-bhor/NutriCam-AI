package validation

import (
	"net/http"
	"server/utils/response"

	"github.com/gin-gonic/gin"
)

type RegistrationRequest struct {
	Mail string  `json:"mail" binding:"required,email"`
	Gid  *string `json:"gid,omitempty"`
	Pass *string `json:"pass,omitempty" binding:"omitempty,min=8"`
}

func RegisterRequestValidation(c *gin.Context) {
	var reqBody RegistrationRequest

	err := c.Bind(&reqBody)

	if err != nil {
		errObj := response.NewErrorResponse(http.StatusBadRequest, err.Error(), "Validation Error")
		response.ErrorResponse(c, errObj)
		return
	}

	c.Set("body", reqBody)
	c.Next()
}
