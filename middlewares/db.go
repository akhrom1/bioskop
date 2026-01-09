package middlewares

import (
	"net/http"

	"bioskop-app/database"

	"github.com/gin-gonic/gin"
)

func CheckDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		if database.DB == nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database belum terkoneksi",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
