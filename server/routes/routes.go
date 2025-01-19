package routes

import (
	authroutes "server/routes/authRoutes"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouterGroups(server *gin.Engine) {
	allRoutes := server.Group("/api")

	authroutes.RegisterAuthRoutes(allRoutes)
}
