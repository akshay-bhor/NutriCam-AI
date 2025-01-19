package authController

import (
	"net/http"
	"server/db"
	"server/middlewares/validation"
	"server/sevices/authService"
	"server/utils/response"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	reqBody, _ := c.Get("body")
	body, ok := reqBody.(validation.RegistrationRequest)

	if !ok {
		response.ErrorResponse(c, http.StatusBadRequest, "", "Failed")
	}
	mail := body.Mail

	// Check if email exists
	var userEmail string
	db.DB.Raw("SELECT mail FROM users WHERE mail = ?", mail).Scan(&userEmail)

	if userEmail != "" {
		response.ErrorResponse(c, http.StatusConflict, "", "Email already exists")
		return
	}

	// Create registration
	userRegistration := authService.RegistrationFactory(c)

	userRegistration.CreateUser()

	response.SuccessResponse(c, "success", "User created successfully")
}
