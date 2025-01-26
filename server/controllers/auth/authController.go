package authController

import (
	"net/http"
	"server/db"
	"server/middlewares/validation"
	"server/sevices/authService"
	"server/sevices/tokenService"
	"server/utils/logger"
	"server/utils/response"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	reqBody, _ := c.Get("body")
	body, ok := reqBody.(validation.RegistrationRequest)

	if !ok {
		errObj := response.NewErrorResponse(400, nil, "Something went wrong")
		response.ErrorResponse(c, errObj)
		return
	}
	mail := body.Mail

	// Check if email exists
	var userEmail string
	db.DB.Raw("SELECT mail FROM users WHERE mail = ?", mail).Scan(&userEmail)

	if userEmail != "" {
		errObj := response.NewErrorResponse(http.StatusConflict, nil, "Email already exists")
		response.ErrorResponse(c, errObj)
		return
	}

	// Create registration
	userRegistration, errObj := authService.RegistrationFactory(c)

	if errObj != nil {
		response.ErrorResponse(c, *errObj)
		return
	}

	user, errObj := userRegistration.CreateUser(c)

	if errObj != nil {
		response.ErrorResponse(c, *errObj)
		return
	}

	userToken := tokenService.UserToken{
		Users: user,
	}

	token, err := userToken.IssueJWTToken()

	if err != nil {
		logger.Error(err, "Error creating user token")
		errObj := response.NewErrorResponse(http.StatusInternalServerError, nil, "Error creating object")
		response.ErrorResponse(c, errObj)
		return
	}

	response.SuccessResponse(c, gin.H{"token": token}, "User created successfully")

}
