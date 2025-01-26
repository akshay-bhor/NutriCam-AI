package routes

import (
	authRoutes "server/routes/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouterGroups(server *gin.Engine) {
	allRouterGroups := server.Group("/api")

	authRoutes.RegisterAuthRoutes(allRouterGroups)
}
