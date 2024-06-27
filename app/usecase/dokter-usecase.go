package usecase

import (
	"klinik/azwan/app/repository"
	"klinik/azwan/entity"
)

type DokterUsecase struct {
	DokterRepo *repository.DokterRepository
}

func NewDokterUsecase(pr *repository.DokterRepository) *DokterUsecase {
	return &DokterUsecase{
		DokterRepo: pr,
	}
}

func (u *DokterUsecase) CreateDokter(dokter *entity.Dokter) error {
	return u.DokterRepo.Create(dokter)
}

func (u *DokterUsecase) GetAllDokters() ([]entity.Dokter, error) {
	return u.DokterRepo.GetAllDokters()
}

func (u *DokterUsecase) GetDokterByID(id int) (entity.Dokter, error) {
	return u.DokterRepo.GetDokterByID(id)
}

func (u *DokterUsecase) UpdateDokter(dokter *entity.Dokter) error {
	return u.DokterRepo.UpdateDokter(dokter)
}

func (u *DokterUsecase) DeleteDokter(id int) error {
	return u.DokterRepo.DeleteDokter(id)
}
