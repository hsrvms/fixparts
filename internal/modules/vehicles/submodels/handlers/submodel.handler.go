package handlers

import (
	"net/http"
	"strconv"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/services"
	"github.com/labstack/echo/v4"
)

type VehicleSubmodelHandler struct {
	service services.VehicleSubmodelService
}

func NewVehicleSubmodelHandler(service services.VehicleSubmodelService) *VehicleSubmodelHandler {
	return &VehicleSubmodelHandler{
		service: service,
	}
}

func (h *VehicleSubmodelHandler) GetAllSubmodels(c echo.Context) error {
	ctx := c.Request().Context()
	submodels, err := h.service.GetAllSubmodels(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, submodels)
}

func (h *VehicleSubmodelHandler) GetSubmodelByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid submodel ID")
	}

	ctx := c.Request().Context()
	submodel, err := h.service.GetSubmodelByID(ctx, id)
	if err != nil {
		if err == vehicleErrors.ErrSubmodelNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, submodel)
}

func (h *VehicleSubmodelHandler) CreateSubmodel(c echo.Context) error {
	submodel := new(models.VehicleSubmodel)
	if err := c.Bind(submodel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.CreateSubmodel(ctx, submodel)
	if err != nil {
		if err == vehicleErrors.ErrModelNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	submodel.SubmodelID = id
	return c.JSON(http.StatusCreated, submodel)
}

func (h *VehicleSubmodelHandler) UpdateSubmodel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid submodel ID")
	}

	submodel := new(models.VehicleSubmodel)
	if err := c.Bind(submodel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	submodel.SubmodelID = id

	ctx := c.Request().Context()
	err = h.service.UpdateSubmodel(ctx, submodel)
	if err != nil {
		switch err {
		case vehicleErrors.ErrSubmodelNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case vehicleErrors.ErrModelNotFound:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, submodel)
}

func (h *VehicleSubmodelHandler) DeleteSubmodel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid submodel ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteSubmodel(ctx, id)
	if err != nil {
		switch err {
		case vehicleErrors.ErrSubmodelNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}
