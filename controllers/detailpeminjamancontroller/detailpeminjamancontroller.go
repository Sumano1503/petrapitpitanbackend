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
	
	if err := models.DB.Where("id_sepeda = ? AND status = ?", id, "on progress").First(&detailPeminjaman).Error; err != nil {
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
	var detailSepedaHalte models.DetailSepedaHalte
	var detailPeminjaman models.DetailPeminjaman
	var cekdDetailPeminjaman []models.DetailPeminjaman
	if err := c.ShouldBindJSON(&detailPeminjaman); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Find(&cekdDetailPeminjaman)

	

	models.DB.Where("id_halte = ? AND status = ?", detailPeminjaman.Id_halte_asal, "available").First(&detailSepedaHalte)
	
	c.JSON(http.StatusOK, gin.H{"detailSepedaHalte": detailSepedaHalte})
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

func HistoryUser(c *gin.Context){
	var detailPeminjaman []models.DetailPeminjaman
	var user models.User

	email := c.Param("email")
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}

	id := user.Id

	if err := models.DB.Where("id_user = ? AND status = ?", id, "done").Find(&detailPeminjaman).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": detailPeminjaman})
}

func DetailHistoryUser(c *gin.Context){
	type DetailHistoryUser struct{
		Id_detail_peminjaman int64 
		Nama_Peminjam string 
		Nrp_Peminjam string 
		Tanggal string 
		Status string 
		Nama_halte_asal string
		Nama_halte_tujuan string 
		Nama_sepeda string 
		Id_sepeda int64 
		Waktu_pengambilan string 
		Waktu_pengembalian string 
		Waktu_Peminjaman string
		Batas_Waktu_Peminjaman string
	}
	var detailhistoryuser DetailHistoryUser
	var detailPeminjaman models.DetailPeminjaman
	var user models.User
	var halteAsal models.Halte
	var halteTujuan models.Halte
	var sepeda models.Sepeda

	id := c.Param("id")
	models.DB.Where("id = ?", id).First(&detailPeminjaman)
	models.DB.Where("id = ?", detailPeminjaman.Id_user).Find(&user)
	models.DB.Where("id_halte = ?", detailPeminjaman.Id_halte_asal).Find(&halteAsal)
	models.DB.Where("id_halte = ?", detailPeminjaman.Id_halte_tujuan).Find(&halteTujuan)
	models.DB.Where("id = ?", detailPeminjaman.Id_sepeda).Find(&sepeda)
	
	detailhistoryuser.Id_detail_peminjaman = detailPeminjaman.Id
	detailhistoryuser.Nama_Peminjam = user.Nama
	detailhistoryuser.Nrp_Peminjam = detailPeminjaman.Nrp_Peminjam
	detailhistoryuser.Tanggal = detailPeminjaman.Tanggal
	detailhistoryuser.Status = detailPeminjaman.Status
	detailhistoryuser.Nama_halte_asal = halteAsal.Nama_halte
	detailhistoryuser.Nama_halte_tujuan = halteTujuan.Nama_halte
	detailhistoryuser.Nama_sepeda = sepeda.Nama
	detailhistoryuser.Id_sepeda = detailPeminjaman.Id_sepeda
	detailhistoryuser.Waktu_pengambilan = detailPeminjaman.Waktu_pengambilan
	detailhistoryuser.Waktu_pengembalian = detailPeminjaman.Waktu_pengembalian
	detailhistoryuser.Waktu_Peminjaman = detailPeminjaman.Waktu_Peminjaman
	detailhistoryuser.Batas_Waktu_Peminjaman = detailPeminjaman.Batas_Waktu_Peminjaman

	c.JSON(http.StatusOK, gin.H{"data": detailhistoryuser})
}