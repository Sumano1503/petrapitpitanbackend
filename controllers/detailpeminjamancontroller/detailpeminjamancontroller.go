package detailpeminjamancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var detailPeminjaman []models.DetailPeminjaman
	models.DB.Find(&detailPeminjaman)
	c.JSON(http.StatusOK, gin.H{"detailPeminjaman": detailPeminjaman})
}

func Show(c *gin.Context) {
	var detailPeminjaman models.DetailPeminjaman
	id := c.Param("id")
	if err := models.DB.First(&detailPeminjaman, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"detailPeminjaman": detailPeminjaman})
}

func ShowIdSep(c *gin.Context) {
	var detailPeminjaman models.DetailPeminjaman
	id := c.Param("id")
	
	if err := models.DB.Where("id_sepeda = ? AND status = ", id, "on progress").First(&detailPeminjaman).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"detailPeminjaman": detailPeminjaman})
}

func Create(c *gin.Context) {
	var detailPeminjaman models.DetailPeminjaman
	if err := c.ShouldBindJSON(&detailPeminjaman); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	models.DB.Create(&detailPeminjaman)
	c.JSON(http.StatusOK, gin.H{"detailPeminjaman": detailPeminjaman})
}

func Update(c *gin.Context) {
	var detailPeminjaman models.DetailPeminjaman
	id := c.Param("id")
	if err := c.ShouldBindJSON(&detailPeminjaman); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}
	if models.DB.Model(&detailPeminjaman).Where("id = ?", id).Updates(&detailPeminjaman).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}
	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var detailPeminjaman models.DetailPeminjaman

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&detailPeminjaman, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}