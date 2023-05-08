package usercontroller

import (
	"net/http"

	"github.com/Sumano1503/petrapitpitanbackend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func Create(c *gin.Context) {
	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	models.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"user": users})
}
func CheckUserSignIn(c *gin.Context){
	var users models.User
	var userData models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	result := models.DB.Where("email = ? ", users.Email).First(&userData)
	if result.Error != nil {
		models.DB.Create(&users)
		c.JSON(http.StatusOK, gin.H{"user": users})
	}else if userData.Role == "Admin" && userData.Nama == "" {
		models.DB.Where("email = ?", users.Email).Updates(&users)
	}else{
		c.JSON(http.StatusOK, gin.H{"user": users, "pesan": "user sudah ada"})
	}
}

func CheckUserAvail(c *gin.Context){
	var userData models.User
	input := c.Param("email")
	

	result := models.DB.Where("email = ? ", input).First(&userData)
	if result.Error != nil{
		c.JSON(http.StatusOK, gin.H{"pesan":"true"})
	}else{
		c.JSON(http.StatusOK, gin.H{"pesan": "false"})
	}
}

func Index(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"user": users})
}

func GetAdmin(c *gin.Context) {
	var users []models.User

	err := models.DB.Where("role = ?", "Admin").Find(&users)

	if err == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}
	c.JSON(http.StatusOK, gin.H{"user": users})
}


func CekAdmin(c *gin.Context) {
	var users models.User

	input := c.Param("email")

	err := models.DB.Where("email = ?", input).Find(&users).Error;

	if err != nil{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	if(users.Role=="Admin"){
		c.JSON(http.StatusOK, gin.H{"user": "User Found Admin"})
	}else if(users.Role == "SuperAdmin"){
		c.JSON(http.StatusCreated, gin.H{"user" : "User Found SuperAdmin"})
	}else{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
	}
	
}

func UserNonAktif(c *gin.Context) {
	var users []models.User
		
	if err := models.DB.Where("status = ?", "nonaktif").Find(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Show(c *gin.Context) {
	var users models.User

	email := c.Param("email")

	if err := models.DB.Where("email = ?", email).Find(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return 
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": users})
	
}



func Update(c *gin.Context) {
	var users models.User

	id := c.Param("id")

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if models.DB.Model(&users).Where("id = ?", id).Updates(&users).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "TIDAK DAPAT MENGUPDATE"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di perbahaarui"})
}

func Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")


	if err := models.DB.Where("id = ?", id ).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user tidak ditemukan"})
        return
	}

	if err:= models.DB.Delete(&user).Error; err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "tidak dapat menghapus"})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"pesan": "berhasil di hapus"})
}


	


