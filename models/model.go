package models

type User struct {
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama string `gorm:"size:255;not null;" json:"nama"`
	Email string `gorm:"size:100;not null;" json:"email"`
	Role string `gorm:"size:100;not null;" json:"role"`
	Status string `gorm:"size:100;not null;" json:"status"`
	Image string `gorm:"size:255;not null;" json:"image"`
	NoTelp string `gorm:"size:255;not null;" json:"no_telp"`
}

type DetailPeminjaman struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama_Peminjam string `gorm:"size:255;not null;" json:"nama_peminjam"`
	Nrp_Peminjam string `gorm:"size:255;not null;" json:"nrp_peminjam"`
	Tanggal string `gorm:"size:255;not null;" json:"tanggal"`
	Id_user int64 `gorm:"size:100;not null;" json:"id_user"`
	Status string `gorm:"size:100;not null;" json:"status"`
	Id_halte_asal int64 `gorm:"size:100;not null;" json:"id_halte_asal"`
	Id_halte_tujuan int64 `gorm:"size:100;not null;" json:"id_halte_tujuan"`
	Id_sepeda int64 `gorm:"size:100;not null;" json:"id_sepeda"`
	Waktu_pengambilan string `gorm:"size:255;not null;" json:"waktu_pengambilan"`
	Waktu_pengembalian string `gorm:"size:255; not null;" json:"waktu_pengembalian"`
	Waktu_Peminjaman string `gorm:"size:255; not null;" json:"waktu_peminjaman"`
	Batas_Waktu_Peminjaman string `gorm:"size:255; not null;" json:"batas_waktu_pengembalian"`
	Sesi int64 `gorm:"size:100;not null;" json:"sesi"`
}

type DetailPelanggaran struct{
	Id_detail_pelanggaran int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Id_pelanggaran int64 `gorm:"size:100;not null;" json:"id_pelanggaran"`
	Id_detail_peminjaman int64 `gorm:"size:100;not null;" json:"id_detail_peminjaman"`
	Id_user int64 `gorm:"size:100;not null;" json:"id_user"`
}

type Pelanggaran struct{
	Id_pelanggaran int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Id_detail_peminjaman int64 `gorm:"size:100;not null;" json:"id_detail_peminjaman"`
	Kode_pelanggaran int64 `gorm:"size:255;not null;" json:"kode_pelanggaran"`
	Id_User int64 `gorm:"size:100;not null;" json:"id_user"`
	Id_Sepeda int64 `gorm:"size:100;not null;" json:"id_sepeda"`
	Tanggal string `gorm:"size:100;not null;" json:"tanggal"`
}

type Halte struct{
	Id_halte int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama_halte string `gorm:"size:255;not null;" json:"nama_halte"`
	Gambar string `gorm:"size:100;not null;" json:"gambar"`
	Tanggal string `gorm:"size:100;not null;" json:"tanggal"`
	Status int `gorm:"size:100;not null;" json:"status"`
}

type DetailSepedaHalte struct{
	Id_detail_sepeda_halte int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Id_halte int64 `gorm:"size:100;not null;" json:"id_halte"`
	Id_sepeda int64 `gorm:"size:100;not null;" json:"id_sepeda"`
	Status string `gorm:"size:100;not null;" json:"status"`
}

type Sepeda struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama string `gorm:"size:255;not null;" json:"nama"`
	Warna string `gorm:"size:100;not null;" json:"warna"`
	Merk string `gorm:"size:100;not null;" json:"merk"`
	Ukuran string `gorm:"size:100;not null;" json:"ukuran"`
	Tipe string `gorm:"size:100;not null;" json:"tipe"`
	Key string `gorm:"size:100;not null;" json:"key"`
	Gambar string `gorm:"size:500;not null;" json:"gambar"`
	Status int64 `gorm:"not null;" json:"status"`
	Tanggal string `gorm:"not null;" json:"tanggal"`
	Alasan string `gorm:"not null;" json:"alasan"`
}

type SesiPeminjaman struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Sesi int64 `gorm:"size:100;not null;" json:"sesi"`
	Waktu_Peminjaman string `gorm:"size:100;not null;" json:"waktu_peminjaman"`
	Batas_Waktu_Penminjaman string `gorm:"size:100;not null;" json:"batas_waktu_peminjaman"`
	Status string `gorm:"size:100;not null;" json:"status"`
}

type Polygon struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Latitude string `gorm:"size:100;not null;" json:"latitude"`
	Longitude string `gorm:"size:100;not null;" json:"longitude"`
}