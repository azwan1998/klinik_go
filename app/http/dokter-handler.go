package http

import (
	"net/http"
	"strconv"

	"klinik/azwan/app/usecase"
	"klinik/azwan/entity"

	"github.com/labstack/echo/v4"
)

type DokterHandler struct {
	DokterUsecase *usecase.DokterUsecase
}

func NewDokterHandler(e *echo.Echo, pu *usecase.DokterUsecase) {
	handler := &DokterHandler{
		DokterUsecase: pu,
	}

	e.GET("/dokters", handler.GetAllDokters)
	e.POST("/dokters", handler.CreateDokter)
	e.GET("/dokters/:id", handler.GetDokterByID)
	e.PUT("/dokters/update/:id", handler.UpdateDokter)
	e.DELETE("/dokters/delete/:id", handler.DeleteDokter)

}

func (h *DokterHandler) CreateDokter(c echo.Context) error {
	var dokter entity.Dokter
	if err := c.Bind(&dokter); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.DokterUsecase.CreateDokter(&dokter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, dokter)
}

func (h *DokterHandler) GetAllDokters(c echo.Context) error {
	dokters, err := h.DokterUsecase.GetAllDokters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dokters)
}

func (h *DokterHandler) GetDokterByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	dokter, err := h.DokterUsecase.GetDokterByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dokter)
}

func (h *DokterHandler) UpdateDokter(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var dokter entity.Dokter
	if err := c.Bind(&dokter); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dokter.ID = id
	err = h.DokterUsecase.UpdateDokter(&dokter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dokter)
}

func (h *DokterHandler) DeleteDokter(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = h.DokterUsecase.DeleteDokter(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
