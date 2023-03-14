package detailsepedahaltecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var detailSepedaHalte []models.DetailSepedaHalte

	models.DB.Find(&detailSepedaHalte)
	c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
}

func Show(c *gin.Context) {
	var detailSepedaHalte models.DetailSepedaHalte

	id := c.Param("id")

	if err := models.DB.First(&detailSepedaHalte, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
	
}

func Create(c *gin.Context) {
	var detailSepedaHalte models.DetailSepedaHalte

	if err := c.ShouldBindJSON(&detailSepedaHalte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&detailSepedaHalte)
	c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
}

func Update(c *gin.Context) {
	var detailSepedaHalte models.DetailSepedaHalte

	id := c.Param("id")

	if err := c.ShouldBindJSON(&detailSepedaHalte); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&detailSepedaHalte).Where("id = ?", id).Updates(&detailSepedaHalte).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var detailSepedaHalte models.DetailSepedaHalte

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&detailSepedaHalte, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}

func DeleteByIdSepeda(c *gin.Context) {	
	var detailSepedaHalte models.DetailSepedaHalte
	id := c.Param("id")


	if err := models.DB.Where("id_sepeda = ?", id ).First(&detailSepedaHalte).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sepeda tidak ditemukan"})
        return
	}

	if err:= models.DB.Delete(&detailSepedaHalte).Error; err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}