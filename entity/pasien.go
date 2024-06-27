package entity

import (
	"time"

	"gorm.io/gorm"
)

type Pasien struct {
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string         `gorm:"type:varchar(255);not null" json:"nama"`
	Umur         int16          `gorm:"not null" json:"umur"`
	JenisKelamin string         `gorm:"type:varchar(255);not null" json:"jenis_kelamin"`
	Alamat       string         `gorm:"type:varchar(255)" json:"alamat"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	NIK          string         `gorm:"type:varchar(255);unique;not null" json:"nik"`
}
