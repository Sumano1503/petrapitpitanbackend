package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpeminjamancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailsepedahaltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/haltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/pelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/sepedacontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/usercontroller"
	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("mysupersecretpharse")

func homepage(w http.ResponseWriter, r *http.Request){
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	fmt.Fprint(w, validToken)
}

func generateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["autorized"] = true
	claims["user"] = "sumano"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequest(){
	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main(){
	handleRequest()

	r := gin.Default();
	models.ConnectDataBase()

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

	
	r.Run()
}