package auth

import (
	"server/db"
	"server/models"
	"server/sevices/tokenService"
	"server/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	// Get the Authorization header
	authHeader := c.GetHeader("Authorization")

	// Set default authentication status
	c.Set("authenticated", false)

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.Next()
		return
	}

	// Extract the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	validToken, userInfo := tokenService.VerifyJWTToken(tokenString)

	if validToken {
		// Fetch user
		var user models.Users
		db.DB.Raw("SELECT * FROM users WHERE id = ?", userInfo.UserId).Scan(&user)

		// Set authentication flag and user info in context
		c.Set("authenticated", true)
		c.Set("user", user)
	}

	c.Next()
}

func RequireAuth(c *gin.Context) {
	authenticated, _ := c.Get("authenticated")

	if !authenticated.(bool) {
		errObj := response.NewErrorResponse(401, nil, "Unauthorized")
		response.ErrorResponse(c, errObj)
		return
	}

	c.Next()
}
