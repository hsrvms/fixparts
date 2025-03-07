package handlers

import (
	"net/http"
	"strconv"

	itemerrors "github.com/hsrvms/fixparts/internal/modules/inventory/items/errors"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/services"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	service        services.ItemService
	barcodeService services.BarcodeService
}

func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{
		service:        service,
		barcodeService: services.NewBarcodeService(),
	}
}

// GetItems handles the retrieval of items with optional filtering
func (h *ItemHandler) GetItems(c echo.Context) error {
	filter := &models.ItemFilter{}

	// Parse query parameters
	if categoryID := c.QueryParam("category_id"); categoryID != "" {
		id, err := strconv.Atoi(categoryID)
		if err == nil {
			filter.CategoryID = &id
		}
	}

	if supplierID := c.QueryParam("supplier_id"); supplierID != "" {
		id, err := strconv.Atoi(supplierID)
		if err == nil {
			filter.SupplierID = &id
		}
	}

	if partNumber := c.QueryParam("part_number"); partNumber != "" {
		filter.PartNumber = &partNumber
	}

	if search := c.QueryParam("search"); search != "" {
		filter.SearchTerm = &search
	}

	if lowStock := c.QueryParam("low_stock"); lowStock == "true" {
		isLowStock := true
		filter.LowStock = &isLowStock
	}

	if isActive := c.QueryParam("is_active"); isActive != "" {
		active := isActive == "true"
		filter.IsActive = &active
	}

	ctx := c.Request().Context()
	items, err := h.service.GetItems(ctx, filter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

// GetLowStockItems handles the retrieval of items with low stock
func (h *ItemHandler) GetLowStockItems(c echo.Context) error {
	ctx := c.Request().Context()
	items, err := h.service.GetLowStockItems(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

// GetItemByID handles the retrieval of a single item by ID
func (h *ItemHandler) GetItemByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item ID")
	}

	ctx := c.Request().Context()
	item, err := h.service.GetItemByID(ctx, id)
	if err != nil {
		if err == itemerrors.ErrItemNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, item)
}

// GetItemByBarcode handles the retrieval of a single item by barcode
func (h *ItemHandler) GetItemByBarcode(c echo.Context) error {
	barcode := c.Param("barcode")
	if barcode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "barcode is required")
	}

	ctx := c.Request().Context()
	item, err := h.service.GetItemByBarcode(ctx, barcode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if item == nil {
		return echo.NewHTTPError(http.StatusNotFound, "item not found")
	}

	return c.JSON(http.StatusOK, item)
}

// CreateItem handles the creation of a new item
func (h *ItemHandler) CreateItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.CreateItem(ctx, item)
	if err != nil {
		switch err {
		case itemerrors.ErrDuplicatePartNumber, itemerrors.ErrDuplicateBarcode:
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	item.ItemID = id
	return c.JSON(http.StatusCreated, item)
}

// UpdateItem handles the update of an existing item
func (h *ItemHandler) UpdateItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item ID")
	}

	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	item.ItemID = id

	ctx := c.Request().Context()
	err = h.service.UpdateItem(ctx, item)
	if err != nil {
		switch err {
		case itemerrors.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case itemerrors.ErrDuplicatePartNumber, itemerrors.ErrDuplicateBarcode:
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, item)
}

// DeleteItem handles the deletion of an item
func (h *ItemHandler) DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid item ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteItem(ctx, id)
	if err != nil {
		switch err {
		case itemerrors.ErrItemNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *ItemHandler) GetBarcodeImage(c echo.Context) error {
	barcode := c.Param("barcode")
	if barcode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "barcode is required")
	}

	imgBytes, err := h.barcodeService.GenerateBarcodeImage(barcode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Blob(http.StatusOK, "image/png", imgBytes)
}
