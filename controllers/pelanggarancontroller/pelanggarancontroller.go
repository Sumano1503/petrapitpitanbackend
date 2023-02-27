package pelanggarancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var pelanggaran []models.Pelanggaran

	models.DB.Find(&pelanggaran)
	c.JSON(http.StatusOK, gin.H{"pelanggaran": pelanggaran})
}

func Show(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	id := c.Param("id")

	if err := models.DB.First(&pelanggaran, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"pelanggaran": pelanggaran})
	
}

func Create(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	if err := c.ShouldBindJSON(&pelanggaran); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&pelanggaran)
	c.JSON(http.StatusOK, gin.H{"pelanggaran": pelanggaran})
}

func Update(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	id := c.Param("id")

	if err := c.ShouldBindJSON(&pelanggaran); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&pelanggaran).Where("id = ?", id).Updates(&pelanggaran).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&pelanggaran, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}