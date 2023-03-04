package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailpeminjamancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/detailsepedahaltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/haltecontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/logincontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/pelanggarancontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/sepedacontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/usercontroller"
	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	IDToken string `json:"id_token"`
}

func ValidateGoogleIDToken(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var token Token
	err := json.NewDecoder(req.Body).Decode(&token)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify Google ID Token
	jwtToken, err := jwt.ParseWithClaims(token.IDToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Replace with your Google OAuth 2.0 client ID
		// https://developers.google.com/identity/protocols/oauth2/web-server#creatingcred
		return []byte("707879399164-esp7a23mv1dnkl6d7asaj6jpdbhtjf37.apps.googleusercontent.com"), nil
	})

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if !jwtToken.Valid {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := jwtToken.Claims.(*jwt.StandardClaims)

	// Verify that the token is for your app
	if claims.Issuer != "https://accounts.google.com" || claims.Audience != "707879399164-esp7a23mv1dnkl6d7asaj6jpdbhtjf37.apps.googleusercontent.com" {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Token is valid, proceed with your business logic here
	fmt.Fprintf(res, "Token is valid")
}

func main(){
	http.HandleFunc("/google/login", logincontroller.GoogleLogin)
	http.HandleFunc("/google/callback", logincontroller.GoogleCallback)
	http.HandleFunc("/google/validate", ValidateGoogleIDToken)
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

	r.Run(":8081")

	
}