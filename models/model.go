package models

type User struct {
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama string `gorm:"size:255;not null;" json:"name"`
	Email string `gorm:"size:100;not null;" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Role string `gorm:"size:100;not null;" json:"role"`
	Status string `gorm:"size:100;not null;" json:"status"`
}

type Sepeda struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama string `gorm:"size:255;not null;" json:"name"`
	Warna string `gorm:"size:100;not null;" json:"warna"`
	Merk string `gorm:"size:100;not null;" json:"merk"`
	Ukuran string `gorm:"size:100;not null;" json:"ukuran"`
	Tipe string `gorm:"size:100;not null;" json:"deskripsi"`
	Gambar string `gorm:"size:100;not null;" json:"gambar"`
}