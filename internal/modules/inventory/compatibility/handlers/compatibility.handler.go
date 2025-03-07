package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	compatibilityerrors "github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/errors"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/services"
	"github.com/labstack/echo/v4"
)

type CompatibilityHandler struct {
	service services.CompatibilityService
}

func NewCompatibilityHandler(service services.CompatibilityService) *CompatibilityHandler {
	return &CompatibilityHandler{
		service: service,
	}
}

func (h *CompatibilityHandler) GetCompatibilities(c echo.Context) error {
	fmt.Print("Hola")
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item ID")
	}

	ctx := c.Request().Context()
	compatibilities, err := h.service.GetCompatibilities(ctx, itemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, compatibilities)
}

func (h *CompatibilityHandler) AddCompatibility(c echo.Context) error {
	compatibility := new(models.Compatibility)
	if err := c.Bind(compatibility); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.AddCompatibility(ctx, compatibility)
	if err != nil {
		switch err {
		case compatibilityerrors.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case compatibilityerrors.ErrCompatibilityExists:
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	compatibility.CompatID = id
	return c.JSON(http.StatusCreated, compatibility)
}

func (h *CompatibilityHandler) RemoveCompatibility(c echo.Context) error {
	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item ID")
	}

	submodelID, err := strconv.Atoi(c.Param("submodelId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid submodel ID")
	}

	ctx := c.Request().Context()
	err = h.service.RemoveCompatibility(ctx, itemID, submodelID)
	if err != nil {
		if err.Error() == "compatibility not found" {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CompatibilityHandler) GetCompatibleItems(c echo.Context) error {
	submodelID, err := strconv.Atoi(c.Param("submodelId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid submodel ID")
	}

	ctx := c.Request().Context()
	items, err := h.service.GetCompatibleItems(ctx, submodelID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}
