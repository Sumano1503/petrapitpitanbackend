package polygoncontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context){
	var polygon []models.Polygon

	models.DB.Find(&polygon)
	c.JSON(http.StatusOK, gin.H{"polygon": polygon})
}

func Show(c *gin.Context){
	var polygon models.Polygon

	id := c.Param("id")

	if err := models.DB.First(&polygon, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"polygon": polygon})
}

func Update(c *gin.Context) {
	var polygon models.Polygon

	id := c.Param("id")

	if err := c.ShouldBindJSON(&polygon); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&polygon).Where("id = ?", id).Updates(&polygon).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Create(c *gin.Context) {
	var polygoninput models.Polygon

	if err := c.ShouldBindJSON(&polygoninput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	models.DB.Create(&polygoninput);
	
	c.JSON(http.StatusOK, gin.H{"polygon": polygoninput})
	
}