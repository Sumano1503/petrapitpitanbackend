package models

type User struct {
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama string `gorm:"size:255;not null;" json:"nama"`
	Email string `gorm:"size:100;not null;" json:"email"`
	Role string `gorm:"size:100;not null;" json:"role"`
	Status string `gorm:"size:100;not null;" json:"status"`
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
	Nama_pelanggaran string `gorm:"size:255;not null;" json:"nama_pelanggaran"`
}

type Halte struct{
	Id_halte int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Nama_halte string `gorm:"size:255;not null;" json:"nama_halte"`
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
}

type SesiPeminjaman struct{
	Id int64 `gorm:"primary_key;auto_increment; unique" json:"id"`
	Sesi int64 `gorm:"size:100;not null;" json:"sesi"`
	Waktu_Peminjaman int64 `gorm:"size:100;not null;" json:"waktu_peminjaman"`
	Batas_Waktu_Penminjaman int64 `gorm:"size:100;not null;" json:"batas_waktu_peminjaman"`
	Status string `gorm:"size:100;not null;" json:"status"`
	Id_Halte int64 `gorm:"size:100;not null;" json:"id_halte"`
}