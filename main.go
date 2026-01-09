package main

import (
	"bioskop-app/database"
	"bioskop-app/middlewares"
	"bioskop-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.Use(middlewares.CheckDB())

	// Route Api Endpoints
	routes.BioskopRoutes(r)

	r.Run(":8080")
}
