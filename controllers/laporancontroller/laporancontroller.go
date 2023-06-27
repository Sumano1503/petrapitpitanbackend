package laporancontroller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func GetLaporan(c *gin.Context){
	type laporan struct{
		BanyakPeminjaman int
		BanyakPelanggaran int
		SepedaBaru int
		HalteBaru int
		TotalSepeda int
		SepedaRusak int
		detailpeminjaman []models.DetailPeminjaman
	}

	type tanggal struct{
		Start string 
		End string
	}

	var pelanggaran []models.Pelanggaran
	var sepeda []models.Sepeda
	var date tanggal
	var listLaporan laporan

	if err := c.ShouldBindJSON(&date); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// getdetailpeminjaman
	queryDetailPeminjaman := fmt.Sprintf("SELECT * FROM detail_peminjamen WHERE STR_TO_DATE(tanggal, '%%d/%%m/%%Y') BETWEEN STR_TO_DATE('%s', '%%d/%%m/%%Y') AND STR_TO_DATE('%s', '%%d/%%m/%%Y')", date.Start, date.End)
	if err := models.DB.Raw(queryDetailPeminjaman).Scan(&listLaporan.detailpeminjaman).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	// getpelanggaran
	queryPelanggaran := fmt.Sprintf("SELECT * FROM pelanggarans WHERE STR_TO_DATE(tanggal, '%%d/%%m/%%Y') BETWEEN STR_TO_DATE('%s', '%%d/%%m/%%Y') AND STR_TO_DATE('%s', '%%d/%%m/%%Y')", date.Start, date.End)
	if err := models.DB.Raw(queryPelanggaran).Scan(&pelanggaran).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	// getsepeda
	models.DB.Find(&sepeda)

	layout := "02/01/2006"
	start, _ := time.Parse(layout, date.Start)
	end, _ := time.Parse(layout, date.End)
	coB:=0
	coR:=0
	for i := 0; i < len(sepeda); i++ {
		tanggal, _ := time.Parse(layout, sepeda[i].Tanggal)

		if tanggal.After(start) && tanggal.Before(end) {
			coB++
		}

		if(sepeda[i].Status == 1){
			coR++;
		}
	}

	listLaporan.BanyakPeminjaman = len(listLaporan.detailpeminjaman)
	listLaporan.BanyakPelanggaran = len(pelanggaran)
	listLaporan.SepedaBaru = coB
	listLaporan.SepedaRusak = coR
	listLaporan.TotalSepeda = len(sepeda)
	
		
	c.JSON(http.StatusOK, gin.H{"data": listLaporan.detailpeminjaman})
}