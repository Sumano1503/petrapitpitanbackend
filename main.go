package main

import (
	"github.com/Sumano1503/petrapitpitanbackend/controllers/sepedacontroller"
	"github.com/Sumano1503/petrapitpitanbackend/controllers/usercontroller"
	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default();
	models.ConnectDataBase()

	r.GET("/api/users", usercontroller.Index)
	r.GET("/api/users/:id", usercontroller.Index)
	r.POST("/api/users", usercontroller.Create)
	r.PUT("/api/users/:id", usercontroller.Update)
	r.DELETE("/api/users", usercontroller.Delete)

	r.GET("/api/sepeda", sepedacontroller.Index)
	r.GET("/api/sepeda/:id", sepedacontroller.Index)
	r.POST("/api/sepeda", sepedacontroller.Create)
	r.PUT("/api/sepeda/:id", sepedacontroller.Update)
	r.DELETE("/api/sepeda", sepedacontroller.Delete)

	r.Run()
}