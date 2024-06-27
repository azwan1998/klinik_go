package entity

import (
	"time"

	"gorm.io/gorm"
)

type Antrian struct {
	ID                int            `gorm:"primaryKey;autoIncrement" json:"id"`
	NomorAntrian      int64          `gorm:"not null" json:"nomor_antrian"`
	Keluhan           string         `gorm:"type:varchar(255);not null" json:"keluhan"`
	PasienID          int            `gorm:"column:id_pasien;not null" json:"id_pasien"`
	DokterID          int            `gorm:"column:id_dokter;not null" json:"id_dokter"`
	TanggalBerkunjung time.Time      `gorm:"not null" json:"tanggal_berkunjung"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type AntrianInput struct {
	NomorAntrian      int64  `json:"nomor_antrian"`
	Keluhan           string `json:"keluhan"`
	PasienID          int    `json:"id_pasien"`
	DokterID          int    `json:"id_dokter"`
	TanggalBerkunjung string `json:"tanggal_berkunjung"`
}

type AntrianDetail struct {
	Antrian
	NamaPasien string `json:"nama_pasien"`
	NamaDokter string `json:"nama_dokter"`
}
