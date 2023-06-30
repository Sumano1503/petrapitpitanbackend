package haltecontroller

import (
	"net/http"
	"time"

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
	waktu:= time.Now()
	halte.Tanggal = waktu.Format("02/01/2006")
	halte.Status = 1
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
	id := c.Param("id")

	if models.DB.Model(&halte).Where("id_halte = ?", id).UpdateColumn("status", 1).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Halte tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "Berhasil dihapus"})
}

func Active(c *gin.Context) {
	var halte models.Halte
	id := c.Param("id")

	if models.DB.Model(&halte).Where("id_halte = ?", id).UpdateColumn("status", 2).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Halte tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "Berhasil dihapus"})
}