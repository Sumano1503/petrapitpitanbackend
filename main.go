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
	"github.com/Sumano1503/petrapitpitanbackend/controllers/sesipeminjamancontroller"
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
			return []byte("w41VrgDgWfr2DxfF6UxYIu7oLoU8rV9YhFzXCdpklE7SmEnN9gWYcdRAduqiMFN"), nil
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

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "berhasil Login"})
	}
}


func main(){
	r := gin.Default();
	models.ConnectDataBase()

	auth := r.Group("/api",AuthMiddleware())

	auth.GET("/login", Login())

	auth.GET("/user", usercontroller.Index)
	auth.GET("/user/:id", usercontroller.Show)
	auth.GET("/cekuser/:email", usercontroller.CekAdmin)
	auth.POST("/user", usercontroller.Create)
	auth.PUT("/user/:id", usercontroller.Update)
	auth.DELETE("/user/:id", usercontroller.Delete)
	auth.GET("/userNonAktif", usercontroller.UserNonAktif)
	auth.POST("/signInCheck", usercontroller.CheckUserSignIn)

	auth.GET("/pelanggaran", pelanggarancontroller.Index)
	auth.GET("/pelanggaran/:id", pelanggarancontroller.Show)
	auth.POST("/pelanggaran", pelanggarancontroller.Create)
	auth.PUT("/pelanggaran/:id", pelanggarancontroller.Update)
	auth.DELETE("/pelanggaran", pelanggarancontroller.Delete)

	auth.GET("/halte", haltecontroller.Index)
	auth.GET("/halte/:id", haltecontroller.Show)
	auth.POST("/halte", haltecontroller.Create)
	auth.PUT("/halte/:id", haltecontroller.Update)
	auth.DELETE("/halte", haltecontroller.Delete)

	auth.GET("/detailsepedahalte", detailsepedahaltecontroller.Index)
	auth.GET("/detailsepedahalte/:id", detailsepedahaltecontroller.Show)
	auth.POST("/detailsepedahalte", detailsepedahaltecontroller.Create)
	auth.PUT("/detailsepedahalte/:id", detailsepedahaltecontroller.Update)
	auth.DELETE("/detailsepedahalte", detailsepedahaltecontroller.Delete)
	auth.DELETE("/detailsepedahalteDelByIdSepeda/:id", detailsepedahaltecontroller.DeleteByIdSepeda)
	auth.GET("/sepedaHalte1", detailsepedahaltecontroller.GetSepedaHalte1)
	auth.GET("/sepedaHalte2", detailsepedahaltecontroller.GetSepedaHalte2)


	auth.GET("/detailpeminjaman", detailpeminjamancontroller.Index)
	auth.GET("/detailpeminjaman/:id", detailpeminjamancontroller.Show)
	auth.GET("/detailpeminjamanbyidsepeda/:id", detailpeminjamancontroller.ShowIdSep)
	auth.POST("/detailpeminjaman", detailpeminjamancontroller.Create)
	auth.PUT("/detailpeminjaman/:id", detailpeminjamancontroller.Update)
	auth.DELETE("/detailpeminjaman", detailpeminjamancontroller.Delete)
	auth.GET("/historyUser/:email", detailpeminjamancontroller.HistoryUser)
	auth.GET("/detailHistoryUser/:id", detailpeminjamancontroller.DetailHistoryUser)

	auth.GET("/detailpelanggaran", detailpelanggarancontroller.Index)
	auth.GET("/detailpelanggaran/:id", detailpelanggarancontroller.Show)
	auth.POST("/detailpelanggaran", detailpelanggarancontroller.Create)
	auth.PUT("/detailpelanggaran/:id", detailpelanggarancontroller.Update)
	auth.DELETE("/detailpelanggaran", detailpelanggarancontroller.Delete)

	auth.GET("/sepeda", sepedacontroller.Index)
	auth.GET("/sepeda/:id", sepedacontroller.Show)
	auth.POST("/sepeda", sepedacontroller.Create)
	auth.PUT("/sepeda/:id", sepedacontroller.Update)
	auth.DELETE("/sepeda/:id", sepedacontroller.Delete)
	
	auth.GET("/sesipeminjaman", sesipeminjamancontroller.GetSesi1Halte1)

	r.Run(":8081")

	
}