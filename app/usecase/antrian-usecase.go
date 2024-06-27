package usecase

import (
	"klinik/azwan/app/repository"
	"klinik/azwan/entity"
)

type AntrianUsecase struct {
	AntrianRepo *repository.AntrianRepository
}

func NewAntrianUsecase(pr *repository.AntrianRepository) *AntrianUsecase {
	return &AntrianUsecase{
		AntrianRepo: pr,
	}
}

func (u *AntrianUsecase) CreateAntrian(antrian *entity.Antrian) error {

	maxNomorAntrian, err := u.AntrianRepo.GetMaxNomorAntrianByDate(antrian.TanggalBerkunjung)
	if err != nil {
		return err
	}
	antrian.NomorAntrian = maxNomorAntrian + 1

	return u.AntrianRepo.Create(antrian)
}

func (u *AntrianUsecase) GetAllAntrians() ([]entity.AntrianDetail, error) {
	return u.AntrianRepo.GetAllAntrians()
}

func (u *AntrianUsecase) UpdateAntrian(antrian *entity.Antrian) error {
	return u.AntrianRepo.UpdateAntrian(antrian)
}

func (u *AntrianUsecase) DeleteAntrian(id int) error {
	return u.AntrianRepo.DeleteAntrian(id)
}

func (au *AntrianUsecase) GetAntrianByID(id int) (*entity.AntrianDetail, error) {
	return au.AntrianRepo.GetAntrianByID(id)
}
