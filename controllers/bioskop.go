package controllers

import (
	"net/http"

	"bioskop-app/database"
	"bioskop-app/models"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var input models.Bioskop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "JSON tidak valid",
		})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan lokasi wajib diisi",
		})
		return
	}

	query := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		input.Nama,
		input.Lokasi,
		input.Rating,
	).Scan(&input.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan data",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    input,
	})
}

func GetBioskop(c *gin.Context) {

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database belum terkoneksi",
		})
		return
	}

	rows, err := database.DB.Query(`
		SELECT id, nama, lokasi, rating
		FROM bioskop
		ORDER BY id ASC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()

	var bioskops []models.Bioskop

	for rows.Next() {
		var b models.Bioskop
		err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		bioskops = append(bioskops, b)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bioskops,
	})
}