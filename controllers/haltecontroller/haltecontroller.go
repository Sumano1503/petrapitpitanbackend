package haltecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var halte []models.Halte

	models.DB.Find(&halte)
	c.JSON(http.StatusOK, gin.H{"halte": halte})
}

func Show(c *gin.Context) {
	var halte models.Halte

	id := c.Param("id")

	if err := models.DB.First(&halte, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"halte": halte})
	
}

func Create(c *gin.Context) {
	var halte models.Halte

	if err := c.ShouldBindJSON(&halte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&halte)
	c.JSON(http.StatusOK, gin.H{"halte": halte})
}

func Update(c *gin.Context) {
	var halte models.Halte

	id := c.Param("id")

	if err := c.ShouldBindJSON(&halte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&halte).Where("id = ?", id).Updates(&halte).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var halte models.Halte

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&halte, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}