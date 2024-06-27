package http

import (
	"net/http"
	"strconv"

	"klinik/azwan/app/usecase"
	"klinik/azwan/entity"

	"github.com/labstack/echo/v4"
)

type PasienHandler struct {
	PasienUsecase *usecase.PasienUsecase
}

func NewPasienHandler(e *echo.Echo, pu *usecase.PasienUsecase) {
	handler := &PasienHandler{
		PasienUsecase: pu,
	}

	e.GET("/pasiens", handler.GetAllPasiens)
	e.GET("/pasiens/:id", handler.GetPasienByID)
	e.POST("/pasiens", handler.CreatePasien)
	e.PUT("/pasiens/update/:id", handler.UpdatePasien)
	e.DELETE("/pasiens/delete/:id", handler.DeletePasien)

}

func (h *PasienHandler) CreatePasien(c echo.Context) error {
	var pasien entity.Pasien
	if err := c.Bind(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.PasienUsecase.CreatePasien(&pasien)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, pasien)
}

func (h *PasienHandler) GetAllPasiens(c echo.Context) error {
	pasiens, err := h.PasienUsecase.GetAllPasiens()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pasiens)
}

func (h *PasienHandler) GetPasienByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	pasien, err := h.PasienUsecase.GetPasienByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pasien)
}

func (h *PasienHandler) UpdatePasien(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var pasien entity.Pasien
	if err := c.Bind(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pasien.ID = id
	err = h.PasienUsecase.UpdatePasien(&pasien)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, pasien)
}

func (h *PasienHandler) DeletePasien(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = h.PasienUsecase.DeletePasien(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
