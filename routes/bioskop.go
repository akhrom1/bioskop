package routes

import (
	"bioskop-app/controllers"

	"github.com/gin-gonic/gin"
)

func BioskopRoutes(r *gin.Engine) {
	bioskop := r.Group("/bioskop")
	{
		bioskop.POST("", controllers.CreateBioskop)
		bioskop.GET("", controllers.GetBioskop)
		bioskop.GET("/:id", controllers.GetBioskopByID)
		bioskop.PUT("/:id", controllers.PutBioskop)
		bioskop.DELETE("/:id", controllers.DeleteBioskop)
	}
}

