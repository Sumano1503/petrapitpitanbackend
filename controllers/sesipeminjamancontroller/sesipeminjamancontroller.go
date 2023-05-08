package sesipeminjamancontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func GetSesi1Halte1(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 1, 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// c.AbortWithStatus(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesiHalte(c *gin.Context){
	var sesiPeminjaman []models.SesiPeminjaman

	var input struct{
		Id_Halte int64 `gorm:"size:100;not null;" json:"id_halte"`
		Sesi int64 `gorm:"size:100;not null;" json:"sesi"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	idHalte:= input.Id_Halte
	idSesi:=input.Sesi

	if result:=models.DB.Where("sesi = ? AND id_halte = ?", idSesi, idHalte).Find(&sesiPeminjaman).Error;result!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error()})
		// c.AbortWithStatus(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sesiPeminjaman})
}

func GetSesi1Halte2(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 1, 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi2Halte1(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 2, 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi2Halte2(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 2, 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi3Halte1(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 3, 1).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
}

func GetSesi3Halte2(c *gin.Context) {
	var sesipeminjaman []models.SesiPeminjaman

	if err := models.DB.Where("sesi = ? AND id_halte = ?", 3, 2).Find(&sesipeminjaman).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sesipeminjaman})
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