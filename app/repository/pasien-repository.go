package repository

import (
	"klinik/azwan/entity"

	"gorm.io/gorm"
)

type PasienRepository struct {
	DB *gorm.DB
}

func NewPasienRepository(db *gorm.DB) *PasienRepository {
	return &PasienRepository{DB: db}
}

func (r *PasienRepository) Create(pasien *entity.Pasien) error {
	return r.DB.Create(pasien).Error
}

func (r *PasienRepository) GetAllPasiens() ([]entity.Pasien, error) {
	var pasiens []entity.Pasien
	err := r.DB.Find(&pasiens).Error
	return pasiens, err
}

func (r *PasienRepository) GetPasienByID(id int) (entity.Pasien, error) {
	var pasien entity.Pasien
	err := r.DB.First(&pasien, id).Error
	return pasien, err
}

func (r *PasienRepository) UpdatePasien(pasien *entity.Pasien) error {
	return r.DB.Save(pasien).Error
}

func (r *PasienRepository) DeletePasien(id int) error {
	return r.DB.Delete(&entity.Pasien{}, id).Error
}
