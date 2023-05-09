package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/petra_pit_pitan"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{} , &Sepeda{}, &Halte{}, &DetailSepedaHalte{}, &Pelanggaran{}, &DetailPelanggaran{}, &DetailPeminjaman{}, &SesiPeminjaman{}, &Polygon{})

	DB = database
}