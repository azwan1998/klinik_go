package http

import (
	"klinik/azwan/app/repository"
	"klinik/azwan/app/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	pasienRepo := repository.NewPasienRepository(db)
	pasienUsecase := usecase.NewPasienUsecase(pasienRepo)
	NewPasienHandler(e, pasienUsecase)

	dokterRepo := repository.NewDokterRepository(db)
	dokterUsecase := usecase.NewDokterUsecase(dokterRepo)
	NewDokterHandler(e, dokterUsecase)

	antrianRepo := repository.NewAntrianRepository(db)
	antrianUsecase := usecase.NewAntrianUsecase(antrianRepo)
	NewAntrianHandler(e, antrianUsecase)

	return e
}
