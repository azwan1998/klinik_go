package repository

import (
	"klinik/azwan/entity"
	"time"

	"gorm.io/gorm"
)

type AntrianRepository struct {
	DB *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) *AntrianRepository {
	return &AntrianRepository{DB: db}
}

func (r *AntrianRepository) Create(antrian *entity.Antrian) error {
	return r.DB.Create(antrian).Error
}

func (r *AntrianRepository) GetAllAntrians() ([]entity.AntrianDetail, error) {
	var antrians []entity.AntrianDetail
	err := r.DB.Table("antrians").
		Select("antrians.*, pasiens.nama as nama_pasien, dokters.nama as nama_dokter").
		Joins("left join pasiens on pasiens.id = antrians.id_pasien").
		Joins("left join dokters on dokters.id = antrians.id_dokter").
		Scan(&antrians).Error
	return antrians, err
}

func (r *AntrianRepository) GetAntrianByID(id int) (*entity.AntrianDetail, error) {
	var antrian entity.AntrianDetail
	err := r.DB.Table("antrians").
		Select("antrians.*, pasiens.nama as nama_pasien, dokters.nama as nama_dokter").
		Joins("left join pasiens on pasiens.id = antrians.id_pasien").
		Joins("left join dokters on dokters.id = antrians.id_dokter").
		Where("antrians.id = ?", id).
		Scan(&antrian).Error
	return &antrian, err
}

func (r *AntrianRepository) UpdateAntrian(antrian *entity.Antrian) error {
	return r.DB.Save(antrian).Error
}

func (r *AntrianRepository) DeleteAntrian(id int) error {
	return r.DB.Delete(&entity.Antrian{}, id).Error
}

func (ar *AntrianRepository) GetMaxNomorAntrianByDate(tanggal time.Time) (int64, error) {
	var maxNomorAntrian int64
	err := ar.DB.Model(&entity.Antrian{}).
		Where("tanggal_berkunjung = ?", tanggal).
		Select("COALESCE(MAX(nomor_antrian), 0)").
		Row().Scan(&maxNomorAntrian)
	return maxNomorAntrian, err
}
