package handlers

import (
	"net/http"
	"strconv"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/services"

	"github.com/labstack/echo/v4"
)

type VehicleModelHandler struct {
	service services.VehicleModelService
}

func NewVehicleModelHandler(service services.VehicleModelService) *VehicleModelHandler {
	return &VehicleModelHandler{
		service: service,
	}
}

func (h *VehicleModelHandler) GetAllModels(c echo.Context) error {
	ctx := c.Request().Context()
	models, err := h.service.GetAllModels(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models)
}

func (h *VehicleModelHandler) GetModelByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid model ID")
	}

	ctx := c.Request().Context()
	model, err := h.service.GetModelByID(ctx, id)
	if err != nil {
		if err == vehicleErrors.ErrModelNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model)
}

func (h *VehicleModelHandler) CreateModel(c echo.Context) error {
	model := new(models.VehicleModel)
	if err := c.Bind(model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.CreateModel(ctx, model)
	if err != nil {
		if err == vehicleErrors.ErrMakeNotFound {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	model.ModelID = id
	return c.JSON(http.StatusCreated, model)
}

func (h *VehicleModelHandler) UpdateModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid model ID")
	}

	model := new(models.VehicleModel)
	if err := c.Bind(model); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	model.ModelID = id

	ctx := c.Request().Context()
	err = h.service.UpdateModel(ctx, model)
	if err != nil {
		switch err {
		case vehicleErrors.ErrModelNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case vehicleErrors.ErrMakeNotFound:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, model)
}

func (h *VehicleModelHandler) DeleteModel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid model ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteModel(ctx, id)
	if err != nil {
		switch err {
		case vehicleErrors.ErrModelNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *VehicleModelHandler) GetSubmodelsByModel(c echo.Context) error {
	modelID, err := strconv.Atoi(c.Param("modelId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid model ID")
	}

	ctx := c.Request().Context()
	submodels, err := h.service.GetSubmodelsByModel(ctx, modelID)
	if err != nil {
		if err == vehicleErrors.ErrModelNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, submodels)
}
