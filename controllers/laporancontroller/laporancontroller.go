package laporancontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLaporan(c *gin.Context){
	type tanggal struct{
		Start string
		End string
	}

	var date tanggal

	if err := c.ShouldBindJSON(&date); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	
	
	c.JSON(http.StatusOK, gin.H{"data": date})
}