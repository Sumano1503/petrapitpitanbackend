package pelanggarancontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var pelanggaran []models.Pelanggaran

	models.DB.Find(&pelanggaran)
	c.JSON(http.StatusOK, gin.H{"pelanggaran": pelanggaran})
}

func Show(c *gin.Context) {
	var pelanggaran []models.Pelanggaran

	id := c.Param("id")

	if err := models.DB.Where("id_user = ?", id).Find(&pelanggaran).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pelanggaran})
	
}

func GetJumlahPelanggaranById(c *gin.Context) {
	var countKode1 int64;
	var countKode2 int64;

	id := c.Param("id")

	if err := models.DB.Where("id_user = ? AND kode_pelanggaran = ?", id, 1).Count(&countKode1).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := models.DB.Where("id_user = ? AND kode_pelanggaran = ?", id, 2).Count(&countKode2).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}


	c.JSON(http.StatusOK, gin.H{"count1": countKode1, "count2" : countKode2})
	
}

func Create(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	if err := c.ShouldBindJSON(&pelanggaran); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&pelanggaran)
	c.JSON(http.StatusOK, gin.H{"pelanggaran": pelanggaran})
}

func Update(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	id := c.Param("id")

	if err := c.ShouldBindJSON(&pelanggaran); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&pelanggaran).Where("id = ?", id).Updates(&pelanggaran).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var pelanggaran models.Pelanggaran

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&pelanggaran, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}