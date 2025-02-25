package main

import (
	"os"
	"server/db"
	"server/migrations"
	"server/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // Auto load env file
	xss "github.com/sahilchopra/gin-gonic-xss-middleware"
)

func init() {
	db.InitDb()
	migrations.Migrate()
}

func main() {
	server := gin.Default()

	// XSS middleware
	var xssMdlwr xss.XssMw
	server.Use(xssMdlwr.RemoveXss())

	routes.RegisterAllRouterGroups(server)

	server.Run(os.Getenv("PORT"))
}
