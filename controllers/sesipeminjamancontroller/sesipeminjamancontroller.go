package sesipeminjamancontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func GetSesi(c *gin.Context){
	var sesiPeminjaman []models.SesiPeminjaman

	idSesi := c.Param("id");

	if result:=models.DB.Where("sesi = ?", idSesi).Find(&sesiPeminjaman).Error;result!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error()})
		// c.AbortWithStatus(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sesiPeminjaman})
}

func GetSesi1(c *gin.Context){
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ?", 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi2(c *gin.Context){
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ?", 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi3(c *gin.Context){
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ?", 3).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func Create(c *gin.Context){
	var sesiPeminjaman models.SesiPeminjaman

	if err := c.ShouldBindJSON(&sesiPeminjaman); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&sesiPeminjaman)
	c.JSON(http.StatusOK, gin.H{"data": sesiPeminjaman})
}

func Delete (c *gin.Context){
	var sesiPeminjaman models.SesiPeminjaman

	id := c.Param("id")

	if err := models.DB.Where("id = ?", id ).First(&sesiPeminjaman).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sepeda tidak ditemukan"})
        return
	}

	if err:= models.DB.Delete(&sesiPeminjaman).Error; err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
} 