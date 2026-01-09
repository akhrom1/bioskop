package controllers

import (
	"database/sql"
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

func GetBioskopByID(c *gin.Context) {

	id := c.Param("id")

	var bioskop models.Bioskop

	err := database.DB.QueryRow(`
		SELECT id, nama, lokasi, rating
		FROM bioskop
		WHERE id = $1
	`, id).Scan(
		&bioskop.ID,
		&bioskop.Nama,
		&bioskop.Lokasi,
		&bioskop.Rating,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Bioskop tidak ditemukan",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bioskop,
	})
}


func PutBioskop(c *gin.Context) {

	id := c.Param("id")

	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan Lokasi tidak boleh kosong",
		})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE bioskop
		SET nama = $1, lokasi = $2, rating = $3
		WHERE id = $4
	`, input.Nama, input.Lokasi, input.Rating, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil diupdate",
	})
}


func DeleteBioskop(c *gin.Context) {

	id := c.Param("id")

	result, err := database.DB.Exec(`
		DELETE FROM bioskop
		WHERE id = $1
	`, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Data bioskop tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bioskop berhasil dihapus",
	})
}
