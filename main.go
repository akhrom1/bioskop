package main

import (
	"bioskop-app/controllers"
	"bioskop-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	// Route Api Endpoints
	r.POST("/bioskop", controllers.CreateBioskop)
	r.GET("/bioskop", controllers.GetBioskop)

	

	r.Run(":8080")
}
