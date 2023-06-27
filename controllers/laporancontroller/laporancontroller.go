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
		HalteTerbanyak string
		TotalSepeda int
		SepedaRusak int
	}

	type tanggal struct{
		Start string 
		End string
	}

	var pelanggaran []models.Pelanggaran
	var detailpeminjaman []models.DetailPeminjaman
	var sepeda []models.Sepeda
	var halte string
	var date tanggal
	var listLaporan laporan

	if err := c.ShouldBindJSON(&date); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	// getdetailpeminjaman
	queryDetailPeminjaman := fmt.Sprintf("SELECT * FROM detail_peminjamen WHERE STR_TO_DATE(tanggal, '%%d/%%m/%%Y') BETWEEN STR_TO_DATE('%s', '%%d/%%m/%%Y') AND STR_TO_DATE('%s', '%%d/%%m/%%Y')", date.Start, date.End)
	if err := models.DB.Raw(queryDetailPeminjaman).Scan(&detailpeminjaman).Error; err != nil {
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

	// gethalte paling sering dikunjungi
	queryHalteTerbanyak := `
	SELECT h.nama_halte 
	FROM detail_peminjamen dp
	JOIN haltes h ON dp.id_halte_tujuan = h.id_halte
	GROUP BY dp.id_halte_tujuan, h.nama_halte
	HAVING COUNT(*) = (
		SELECT MAX(Jumlah)
		FROM (
			SELECT id_halte_tujuan, COUNT(*) AS Jumlah
			FROM detail_peminjamen
			GROUP BY id_halte_tujuan
		) AS Counts
	);
`

if err := models.DB.Raw(queryHalteTerbanyak).Scan(&halte).Error; err != nil {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	return
}


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




	listLaporan.BanyakPeminjaman = len(detailpeminjaman)
	listLaporan.BanyakPelanggaran = len(pelanggaran)
	listLaporan.SepedaBaru = coB
	listLaporan.SepedaRusak = coR
	listLaporan.TotalSepeda = len(sepeda)
	listLaporan.HalteTerbanyak = halte
	
		
	c.JSON(http.StatusOK, gin.H{"detailPeminjaman": detailpeminjaman, "laporan": listLaporan})
}