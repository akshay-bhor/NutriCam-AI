package authRoutes

import (
	authController "server/controllers/auth"
	authMiddleware "server/middlewares/auth"
	"server/middlewares/validation"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(routerGroup *gin.RouterGroup) {
	auth := routerGroup.Group("/auth")

	auth.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hi": "hello world",
		})
	})

	auth.POST("/register", validation.RegisterRequestValidation, authController.RegisterUser)

	auth.GET("/totp", authMiddleware.IsAuthenticated, authMiddleware.RequireAuth, authController.SetupTopt)
}
