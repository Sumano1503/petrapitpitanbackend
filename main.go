package main

import (
	"encoding/base64"
	"encoding/json"
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
	"github.com/golang-jwt/jwt"
)

  func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		claims := jwt.MapClaims{}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tkn := strings.Split(tokenString, ".")
		headerBytes, err := base64.RawURLEncoding.DecodeString(tkn[0])
		if err != nil {
			return
		}
		err = json.Unmarshal(headerBytes, gin.H{"error" : err})
		if err != nil {
			return
		}
		// Decode payload
		payloadBytes, err := base64.RawURLEncoding.DecodeString(tkn[1])
		if err != nil {
			return
		}
		err = json.Unmarshal(payloadBytes, gin.H{"error" : err})
		if err != nil {
			return
		}
		signatureBytes, err:= base64.RawURLEncoding.DecodeString(tkn[2])
		err = json.Unmarshal(signatureBytes, gin.H{"error" : err})
		decodedToken := string(headerBytes) + "." + string(payloadBytes) + "." + string(signatureBytes) 
		c.JSON(http.StatusUnauthorized, gin.H{"token": decodedToken})
		c.Abort()
		token, err := jwt.ParseWithClaims(decodedToken, claims, func(token *jwt.Token) (interface{}, error) {
			// Replace this with your own key lookup logic
			return []byte(""), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}


func main(){
	r := gin.Default();
	models.ConnectDataBase()

	auth := r.Group("/api",AuthMiddleware())
	auth.GET("/user", usercontroller.Index)
	auth.GET("/user/:id", usercontroller.Show)
	auth.POST("/user", usercontroller.Create)
	auth.PUT("/user/:id", usercontroller.Update)
	auth.DELETE("/user", usercontroller.Delete)

	auth.GET("/pelanggaran", pelanggarancontroller.Index)
	auth.GET("/pelanggaran/:id", pelanggarancontroller.Show)
	auth.POST("/pelanggaran", pelanggarancontroller.Create)
	auth.PUT("/pelanggaran/:id", pelanggarancontroller.Update)
	auth.DELETE("/pelanggaran", pelanggarancontroller.Delete)

	auth.GET("/api/halte", haltecontroller.Index)
	auth.GET("/api/halte/:id", haltecontroller.Show)
	auth.POST("/api/halte", haltecontroller.Create)
	auth.PUT("/api/halte/:id", haltecontroller.Update)
	auth.DELETE("/api/halte", haltecontroller.Delete)

	auth.GET("/api/detailsepedahalte", detailsepedahaltecontroller.Index)
	auth.GET("/api/detailsepedahalte/:id", detailsepedahaltecontroller.Show)
	auth.POST("/api/detailsepedahalte", detailsepedahaltecontroller.Create)
	auth.PUT("/api/detailsepedahalte/:id", detailsepedahaltecontroller.Update)
	auth.DELETE("/api/detailsepedahalte", detailsepedahaltecontroller.Delete)

	auth.GET("/api/detailpeminjaman", detailpeminjamancontroller.Index)
	auth.GET("/api/detailpeminjaman/:id", detailpeminjamancontroller.Show)
	auth.POST("/api/detailpeminjaman", detailpeminjamancontroller.Create)
	auth.PUT("/api/detailpeminjaman/:id", detailpeminjamancontroller.Update)
	auth.DELETE("/api/detailpeminjaman", detailpeminjamancontroller.Delete)

	auth.GET("/api/detailpelanggaran", detailpelanggarancontroller.Index)
	auth.GET("/api/detailpelanggaran/:id", detailpelanggarancontroller.Show)
	auth.POST("/api/detailpelanggaran", detailpelanggarancontroller.Create)
	auth.PUT("/api/detailpelanggaran/:id", detailpelanggarancontroller.Update)
	auth.DELETE("/api/detailpelanggaran", detailpelanggarancontroller.Delete)

	auth.GET("/sepeda", sepedacontroller.Index)
	auth.GET("/api/sepeda/:id", sepedacontroller.Show)
	auth.POST("/api/sepeda", sepedacontroller.Create)
	auth.PUT("/api/sepeda/:id", sepedacontroller.Update)
	auth.DELETE("/api/sepeda", sepedacontroller.Delete)
	
	r.Run(":8081")

	
}