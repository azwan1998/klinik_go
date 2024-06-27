package repository

import (
	"klinik/azwan/entity"

	"gorm.io/gorm"
)

type DokterRepository struct {
	DB *gorm.DB
}

func NewDokterRepository(db *gorm.DB) *DokterRepository {
	return &DokterRepository{DB: db}
}

func (r *DokterRepository) Create(dokter *entity.Dokter) error {
	return r.DB.Create(dokter).Error
}

func (r *DokterRepository) GetAllDokters() ([]entity.Dokter, error) {
	var dokters []entity.Dokter
	err := r.DB.Find(&dokters).Error
	return dokters, err
}

func (r *DokterRepository) GetDokterByID(id int) (entity.Dokter, error) {
	var dokter entity.Dokter
	err := r.DB.First(&dokter, id).Error
	return dokter, err
}

func (r *DokterRepository) UpdateDokter(dokter *entity.Dokter) error {
	return r.DB.Save(dokter).Error
}

func (r *DokterRepository) DeleteDokter(id int) error {
	return r.DB.Delete(&entity.Dokter{}, id).Error
}
