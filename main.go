package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpeminjamancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailsepedahaltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/haltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/pelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/sepedacontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/usercontroller"
	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")[1] // mengambil token setelah "Bearer "
	// gunakan token untuk verifikasi pengguna

	fmt.Println(token)
  }

func main(){
	r := gin.Default();
	models.ConnectDataBase()

	http.HandleFunc("/api/TokenAuth", handler)

	r.GET("/api/user", usercontroller.Index)
	r.GET("/api/user/:id", usercontroller.Show)
	r.POST("/api/user", usercontroller.Create)
	r.PUT("/api/user/:id", usercontroller.Update)
	r.DELETE("/api/user", usercontroller.Delete)

	r.GET("/api/pelanggaran", pelanggarancontroller.Index)
	r.GET("/api/pelanggaran/:id", pelanggarancontroller.Show)
	r.POST("/api/pelanggaran", pelanggarancontroller.Create)
	r.PUT("/api/pelanggaran/:id", pelanggarancontroller.Update)
	r.DELETE("/api/pelanggaran", pelanggarancontroller.Delete)

	r.GET("/api/halte", haltecontroller.Index)
	r.GET("/api/halte/:id", haltecontroller.Show)
	r.POST("/api/halte", haltecontroller.Create)
	r.PUT("/api/halte/:id", haltecontroller.Update)
	r.DELETE("/api/halte", haltecontroller.Delete)

	r.GET("/api/detailsepedahalte", detailsepedahaltecontroller.Index)
	r.GET("/api/detailsepedahalte/:id", detailsepedahaltecontroller.Show)
	r.POST("/api/detailsepedahalte", detailsepedahaltecontroller.Create)
	r.PUT("/api/detailsepedahalte/:id", detailsepedahaltecontroller.Update)
	r.DELETE("/api/detailsepedahalte", detailsepedahaltecontroller.Delete)

	r.GET("/api/detailpeminjaman", detailpeminjamancontroller.Index)
	r.GET("/api/detailpeminjaman/:id", detailpeminjamancontroller.Show)
	r.POST("/api/detailpeminjaman", detailpeminjamancontroller.Create)
	r.PUT("/api/detailpeminjaman/:id", detailpeminjamancontroller.Update)
	r.DELETE("/api/detailpeminjaman", detailpeminjamancontroller.Delete)

	r.GET("/api/detailpelanggaran", detailpelanggarancontroller.Index)
	r.GET("/api/detailpelanggaran/:id", detailpelanggarancontroller.Show)
	r.POST("/api/detailpelanggaran", detailpelanggarancontroller.Create)
	r.PUT("/api/detailpelanggaran/:id", detailpelanggarancontroller.Update)
	r.DELETE("/api/detailpelanggaran", detailpelanggarancontroller.Delete)

	r.GET("/api/sepeda", sepedacontroller.Index)
	r.GET("/api/sepeda/:id", sepedacontroller.Show)
	r.POST("/api/sepeda", sepedacontroller.Create)
	r.PUT("/api/sepeda/:id", sepedacontroller.Update)
	r.DELETE("/api/sepeda", sepedacontroller.Delete)

	r.Run(":8081")

	
}