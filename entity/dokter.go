package entity

import (
	"time"

	"gorm.io/gorm"
)

type Dokter struct {
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama         string         `gorm:"type:varchar(255);not null" json:"nama"`
	JenisKelamin string         `gorm:"type:varchar(255);not null" json:"jenis_kelamin"`
	Umur         int            `gorm:"not null" json:"umur"`
	Poli         string         `gorm:"type:varchar(255);not null" json:"poli"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
