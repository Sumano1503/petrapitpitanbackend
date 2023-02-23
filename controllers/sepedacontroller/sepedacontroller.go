package sepedacontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var sepeda []models.Sepeda

	models.DB.Find(&sepeda)
	c.JSON(http.StatusOK, gin.H{"Sepeda": sepeda})
}

func Create(c *gin.Context) {
	
}

func Update(c *gin.Context) {
	
}

func Delete(c *gin.Context) {
	
}

func Show(c *gin.Context) {
	
}