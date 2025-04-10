package http

import (
	"errors"
	"net/http"
	"strconv"

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
	e.GET("/antrians/query/", handler.SearchingAntrian)
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

	antrian, err := h.AntrianUsecase.CreateAntrian(input)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidDateFormat) {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
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

func (h *AntrianHandler) SearchingAntrian(c echo.Context) error {
	searchQuery := c.QueryParam("searching")
	if searchQuery == "" {
		return c.JSON(http.StatusBadRequest, "searching query not valid")
	}

	antrians, err := h.AntrianUsecase.SearchingAntrian(searchQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, antrians)
}

func (h *AntrianHandler) UpdateAntrian(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid ID")
	}

	var input entity.AntrianInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	antrian, err := h.AntrianUsecase.UpdateAntrian(id, input)
	if err != nil {
		if err == usecase.ErrAntrianNotFound {
			return c.JSON(http.StatusNotFound, "antrian not found")
		}
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
