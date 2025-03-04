package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// InventoryHandler handles inventory-related requests
type InventoryHandler struct {
	// For testing purposes, we'll just use a hardcoded value
	lowStockCount int
}

// NewInventoryHandler creates a new instance of InventoryHandler
func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{
		lowStockCount: 5, // Hardcoded value for testing
	}
}

// GetLowStockCount handles the GET request for /api/inventory/low-stock-count
func (h *InventoryHandler) GetLowStockCount(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%d", h.lowStockCount))
	// return c.String(http.StatusOK, fmt.Sprint("-"))
}
