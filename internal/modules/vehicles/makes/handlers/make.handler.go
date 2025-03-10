package handlers

import (
	"net/http"
	"strconv"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/services"
	"github.com/labstack/echo/v4"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
)

type VehicleMakeHandler struct {
	service services.VehicleMakeService
}

func NewVehicleMakeHandler(service services.VehicleMakeService) *VehicleMakeHandler {
	return &VehicleMakeHandler{
		service: service,
	}
}

// Make handlers
func (h *VehicleMakeHandler) GetAllMakes(c echo.Context) error {
	ctx := c.Request().Context()
	makes, err := h.service.GetAllMakes(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, makes)
}

func (h *VehicleMakeHandler) GetMakeByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid make ID")
	}

	ctx := c.Request().Context()
	make, err := h.service.GetMakeByID(ctx, id)
	if err != nil {
		if err == vehicleErrors.ErrMakeNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, make)
}

func (h *VehicleMakeHandler) CreateMake(c echo.Context) error {
	make := new(models.VehicleMake)
	if err := c.Bind(make); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.CreateMake(ctx, make)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	make.MakeID = id
	return c.JSON(http.StatusCreated, make)
}

func (h *VehicleMakeHandler) UpdateMake(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid make ID")
	}

	make := new(models.VehicleMake)
	if err := c.Bind(make); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	make.MakeID = id

	ctx := c.Request().Context()
	err = h.service.UpdateMake(ctx, make)
	if err != nil {
		switch err {
		case vehicleErrors.ErrMakeNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, make)
}

func (h *VehicleMakeHandler) DeleteMake(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid make ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteMake(ctx, id)
	if err != nil {
		switch err {
		case vehicleErrors.ErrMakeNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *VehicleMakeHandler) GetModelsByMake(c echo.Context) error {
	makeID, err := strconv.Atoi(c.Param("makeId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid make ID")
	}

	ctx := c.Request().Context()
	models, err := h.service.GetModelsByMake(ctx, makeID)
	if err != nil {
		if err == vehicleErrors.ErrMakeNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, models)
}
