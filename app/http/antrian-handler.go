package http

import (
	"net/http"
	"strconv"
	"time"

	"klinik/azwan/app/usecase"
	"klinik/azwan/entity"

	"github.com/labstack/echo/v4"
)

type AntrianHandler struct {
	AntrianUsecase *usecase.AntrianUsecase
}

func NewAntrianHandler(e *echo.Echo, pu *usecase.AntrianUsecase) {
	handler := &AntrianHandler{
		AntrianUsecase: pu,
	}

	e.GET("/antrians", handler.GetAllAntrians)
	e.POST("/antrians", handler.CreateAntrian)
	e.GET("/antrians/:id", handler.GetAntrianByID)
	e.PUT("/antrians/update/:id", handler.UpdateAntrian)
	e.DELETE("/antrians/delete/:id", handler.DeleteAntrian)

}

func (h *AntrianHandler) CreateAntrian(c echo.Context) error {
	var input entity.AntrianInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tanggal, err := time.Parse("2006-01-02", input.TanggalBerkunjung)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid date format")
	}

	antrian := entity.Antrian{
		Keluhan:           input.Keluhan,
		PasienID:          input.PasienID,
		DokterID:          input.DokterID,
		TanggalBerkunjung: tanggal,
	}

	err = h.AntrianUsecase.CreateAntrian(&antrian)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, antrian)
}

func (h *AntrianHandler) GetAllAntrians(c echo.Context) error {
	antrians, err := h.AntrianUsecase.GetAllAntrians()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, antrians)
}

func (h *AntrianHandler) GetAntrianByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	antrians, err := h.AntrianUsecase.GetAntrianByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, antrians)
}

func (h *AntrianHandler) UpdateAntrian(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var input entity.AntrianInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tanggal, err := time.Parse("2006-01-02", input.TanggalBerkunjung)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid date format")
	}

	getAntrian, err := h.AntrianUsecase.GetAntrianByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Antrian not found")
	}

	antrian := &entity.Antrian{
		ID:                getAntrian.ID,
		NomorAntrian:      getAntrian.NomorAntrian,
		Keluhan:           input.Keluhan,
		PasienID:          input.PasienID,
		DokterID:          input.DokterID,
		TanggalBerkunjung: tanggal,
	}

	err = h.AntrianUsecase.UpdateAntrian(antrian)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, antrian)
}

func (h *AntrianHandler) DeleteAntrian(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = h.AntrianUsecase.DeleteAntrian(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
