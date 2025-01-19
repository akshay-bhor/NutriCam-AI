package authroutes

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(server *gin.RouterGroup) {
	auth := server.Group("/auth")

	auth.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hi": "hello",
		})
	})
}
