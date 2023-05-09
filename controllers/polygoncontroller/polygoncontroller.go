package polygoncontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	var polygon []models.Polygon

	models.DB.Find(&polygon)
	c.JSON(http.StatusOK, gin.H{"polygon": polygon})
}

func Create(c *gin.Context) {
	var polygon models.Polygon

	if err := c.ShouldBindJSON(&polygon); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	if models.DB.Model(&polygon).Where("id = ?", polygon.Id).RowsAffected == 0{
		models.DB.Create(&polygon)
	}else{
		if models.DB.Model(&polygon).Where("id = ?", polygon.Id).Updates(&polygon).RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
			return 
		}
	}
	
	
	c.JSON(http.StatusOK, gin.H{"polygon": polygon})
}