package sesipeminjamancontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func GetSesi1Halte1(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 1, 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		// c.AbortWithStatus(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi1Halte2(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 1, 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi2Halte1(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 2, 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi2Halte2(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 2, 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}