package laporancontroller

import (
	"fmt"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func GetLaporan(c *gin.Context){
	type tanggal struct{
		Start string 
		End string
	}

	var detail_peminjamen []models.DetailPeminjaman

	var date tanggal

	if err := c.ShouldBindJSON(&date); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	query := fmt.Sprintf("SELECT * FROM detail_peminjaman WHERE STR_TO_DATE(tanggal, '%%d/%%m/%%Y') BETWEEN STR_TO_DATE('%s', '%%d/%%m/%%Y') AND STR_TO_DATE('%s', '%%d/%%m/%%Y')", date.Start, date.End)

	// Menjalankan query dan mendapatkan hasil
	if err := models.DB.Raw(query).Scan(&detail_peminjamen).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	
	
	c.JSON(http.StatusOK, gin.H{"data": detail_peminjamen})
}