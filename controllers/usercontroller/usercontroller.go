package usercontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"user": users})
}

func Create(c *gin.Context) {
	
}

func Update(c *gin.Context) {
	
}

func Delete(c *gin.Context) {
	
}

func Show(c *gin.Context) {

}