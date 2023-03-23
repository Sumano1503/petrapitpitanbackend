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
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}