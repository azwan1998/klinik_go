package usecase

import (
	"errors"
	"klinik/azwan/app/repository"
	"klinik/azwan/entity"
	"time"
)

var ErrAntrianNotFound = errors.New("antrian not found")

var ErrInvalidDateFormat = errors.New("invalid date format")

type AntrianUsecase struct {
	AntrianRepo *repository.AntrianRepository
}

func NewAntrianUsecase(pr *repository.AntrianRepository) *AntrianUsecase {
	return &AntrianUsecase{
		AntrianRepo: pr,
	}
}

func (u *AntrianUsecase) CreateAntrian(input entity.AntrianInput) (*entity.Antrian, error) {
	tanggal, err := time.Parse("2006-01-02", input.TanggalBerkunjung)
	if err != nil {
		return nil, ErrInvalidDateFormat
	}

	lastAntrian, err := u.AntrianRepo.GetMaxNomorAntrianByDate(tanggal)
	if err != nil {
		return nil, err
	}

	var nomorAntrian int64 = 1
	if lastAntrian > 0 {
		nomorAntrian = lastAntrian + 1
	}

	antrian := &entity.Antrian{
		NomorAntrian:      nomorAntrian,
		Keluhan:           input.Keluhan,
		PasienID:          input.PasienID,
		DokterID:          input.DokterID,
		TanggalBerkunjung: tanggal,
	}

	if err := u.AntrianRepo.Create(antrian); err != nil {
		return nil, err
	}

	return antrian, nil
}

func (u *AntrianUsecase) GetAllAntrians() ([]entity.AntrianDetail, error) {
	return u.AntrianRepo.GetAllAntrians()
}

func (u *AntrianUsecase) UpdateAntrian(id int, input entity.AntrianInput) (*entity.Antrian, error) {
	getAntrianDetail, err := u.AntrianRepo.GetAntrianByID(id)
	if err != nil {
		return nil, ErrAntrianNotFound
	}

	tanggal, err := time.Parse("2006-01-02", input.TanggalBerkunjung)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	getAntrian := &entity.Antrian{
		ID:                getAntrianDetail.ID,
		NomorAntrian:      getAntrianDetail.NomorAntrian,
		Keluhan:           input.Keluhan,
		PasienID:          input.PasienID,
		DokterID:          input.DokterID,
		TanggalBerkunjung: tanggal,
	}

	if err := u.AntrianRepo.UpdateAntrian(getAntrian); err != nil {
		return nil, err
	}

	return getAntrian, nil
}

func (u *AntrianUsecase) DeleteAntrian(id int) error {
	return u.AntrianRepo.DeleteAntrian(id)
}

func (u *AntrianUsecase) GetAntrianByID(id int) (*entity.AntrianDetail, error) {
	return u.AntrianRepo.GetAntrianByID(id)
}

func (u *AntrianUsecase) SearchingAntrian(searchQuery string) ([]entity.AntrianDetail, error) {
	return u.AntrianRepo.SearchingAntrian(searchQuery)
}
