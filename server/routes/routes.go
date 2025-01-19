package routes

import (
	"server/routes/authroutess"

	"github.com/gin-gonic/gin"
)

func RegisterAllRouterGroups(server *gin.Engine) {
	allRoutes := server.Group("/api")

	authroutess.RegisterAuthRoutes(allRoutes)
}
