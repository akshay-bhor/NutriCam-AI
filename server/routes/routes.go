package routes

import (
	"server/routes/authroutes"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouterGroups(server *gin.Engine) {
	allRoutes := server.Group("/api")

	authroutes.RegisterAuthRoutes(allRoutes)
}
