package usecase

import (
	"klinik/azwan/app/repository"
	"klinik/azwan/entity"
)

type PasienUsecase struct {
	PasienRepo *repository.PasienRepository
}

func NewPasienUsecase(pr *repository.PasienRepository) *PasienUsecase {
	return &PasienUsecase{
		PasienRepo: pr,
	}
}

func (u *PasienUsecase) CreatePasien(pasien *entity.Pasien) error {
	return u.PasienRepo.Create(pasien)
}

func (u *PasienUsecase) GetAllPasiens() ([]entity.Pasien, error) {
	return u.PasienRepo.GetAllPasiens()
}

func (u *PasienUsecase) GetPasienByID(id int) (entity.Pasien, error) {
	return u.PasienRepo.GetPasienByID(id)
}

func (u *PasienUsecase) UpdatePasien(pasien *entity.Pasien) error {
	return u.PasienRepo.UpdatePasien(pasien)
}

func (u *PasienUsecase) DeletePasien(id int) error {
	return u.PasienRepo.DeletePasien(id)
}
