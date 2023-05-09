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
	var polygoninput models.Polygon

	if err := c.ShouldBindJSON(&polygoninput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	result := models.DB.Table("polygons").Where("id = ? ", polygoninput.Id)
	if result.Error != nil{
		if models.DB.Model(&polygoninput).Where("id = ?", polygoninput.Id).Updates(&polygoninput).RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
			return 
		}
	}else{
		models.DB.Create(&polygoninput);
	}
	c.JSON(http.StatusOK, gin.H{"polygon": polygoninput})
}