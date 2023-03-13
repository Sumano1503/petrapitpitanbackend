package sepedacontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var sepeda []models.Sepeda

	models.DB.Find(&sepeda)
	c.JSON(http.StatusOK, gin.H{"data": sepeda})
}

func Show(c *gin.Context) {
	var sepeda models.Sepeda

	id := c.Param("id")

	if err := models.DB.First(&sepeda, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": sepeda})
	
}

func Create(c *gin.Context) {
	var sepeda models.Sepeda

	if err := c.ShouldBindJSON(&sepeda); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&sepeda)
	c.JSON(http.StatusOK, gin.H{"data": sepeda})
}

func Update(c *gin.Context) {
	var sepeda models.Sepeda

	id := c.Param("id")

	if err := c.ShouldBindJSON(&sepeda); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&sepeda).Where("id = ?", id).Updates(&sepeda).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var sepeda models.Sepeda
	id := c.Param("id")


	if err := c.ShouldBindJSON(&sepeda); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Delete(&sepeda).Where("id = ?", id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}

