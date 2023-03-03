package main

import (
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
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func generateToken() (string, error) {
    // Set claims
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["user_id"] = "123"
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour

    // Generate token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := []byte("your-secret-key")
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
    // Create OAuth2 configuration
    oauthConfig := &oauth2.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret",
        RedirectURL:  "your-redirect-url",
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }

    // Get authorization URL
    authURL := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

    // Redirect user to authorization URL
    http.Redirect(w, r, authURL, http.StatusSeeOther)
}

func main(){
	r := gin.Default();
	models.ConnectDataBase()

	http.HandleFunc("/google/login", googleLoginHandler)

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

	r.Run(":5001")

	
}