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

func GetSepedaHalte1(c *gin.Context) {
	var detailSepedaHalte []models.DetailSepedaHalte;
	var sepeda []models.Sepeda;
	var sepedaHalte1 []models.Sepeda;
	models.DB.Find(&detailSepedaHalte)
		// c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
	models.DB.Find(&sepeda)
		// c.JSON(http.StatusOK, gin.H{"sepeda": sepeda})
	models.DB.Where("id_halte = ? AND status = ?", 1, "available").Find(&detailSepedaHalte)
	for i:= 0; i < len(detailSepedaHalte); i++ {
		for j:= 0; j < len(sepeda); j++ {
			if detailSepedaHalte[i].Id_sepeda == sepeda[j].Id {
				sepedaHalte1 = append(sepedaHalte1, sepeda[j])
			}
		}
	}
	
	
	c.JSON(http.StatusOK, gin.H{"sepedaHalte1": sepedaHalte1})
}

func GetSepedaHalte2(c *gin.Context) {
	var detailSepedaHalte []models.DetailSepedaHalte;
	var sepeda []models.Sepeda;
	var sepedaHalte2 []models.Sepeda;
	models.DB.Find(&detailSepedaHalte)
		// c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
	models.DB.Find(&sepeda)
		// c.JSON(http.StatusOK, gin.H{"sepeda": sepeda})
	models.DB.Where("id_halte = ? AND status = ?", 2, "available").Find(&detailSepedaHalte)
	for i:= 0; i < len(detailSepedaHalte); i++ {
		for j:= 0; j < len(sepeda); j++ {
			if detailSepedaHalte[i].Id_sepeda == sepeda[j].Id {
				sepedaHalte2 = append(sepedaHalte2, sepeda[j])
			}
		}
	}
	
	
	c.JSON(http.StatusOK, gin.H{"sepedaHalte2": sepedaHalte2})
}

func GetSepedaDipinjam(c *gin.Context) {
	var detailSepedaHalte []models.DetailSepedaHalte;
	var sepeda []models.Sepeda;
	var sepedaDipinjam []models.Sepeda;
	models.DB.Find(&detailSepedaHalte)
		// c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
	models.DB.Find(&sepeda)
		// c.JSON(http.StatusOK, gin.H{"sepeda": sepeda})
	models.DB.Where("status = ?", "dipinjam").Find(&detailSepedaHalte)
	for i:= 0; i < len(detailSepedaHalte); i++ {
		for j:= 0; j < len(sepeda); j++ {
			if detailSepedaHalte[i].Id_sepeda == sepeda[j].Id {
				sepedaDipinjam = append(sepedaDipinjam, sepeda[j])
			}
		}
	}
	
	
	c.JSON(http.StatusOK, gin.H{"sepedaDipinjam": sepedaDipinjam})
}

func CekSepedaTerpakai(c *gin.Context){
	var detailSepedaHalte []models.DetailSepedaHalte;

	idHalte := c.Param("idHalte")

	models.DB.Where("id_halte = ? AND status = ?", idHalte, "available").First(&detailSepedaHalte)

	if(len(detailSepedaHalte)<1){
		c.JSON(http.StatusBadGateway, gin.H{"pesan": "Semua Sepeda Terpakai"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"pesan": "Ada Sepeda Yang Tersedia ", "detailSepedaHalte": detailSepedaHalte})
		return
	}
}