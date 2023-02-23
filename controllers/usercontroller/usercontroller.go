package usercontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Show(c *gin.Context) {
	var users models.User

	id := c.Param("id")

	if err := models.DB.First(&users, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
	
}

func Create(c *gin.Context) {
	
	
}

func Update(c *gin.Context) {
	
}

func Delete(c *gin.Context) {
	
}

