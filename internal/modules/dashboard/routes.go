package dashboard

import (
	"github.com/hsrvms/fixparts/internal/modules/dashboard/handlers"
	"github.com/hsrvms/fixparts/internal/modules/dashboard/repositories"
	"github.com/hsrvms/fixparts/internal/modules/dashboard/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {

	repo := repositories.NewPostgresDashboardRepository(database)
	service := services.NewDashboardService(repo)
	apiHandler := handlers.NewDashboardAPIHandler(service)

	e.GET("/", handlers.ViewHandler)

	// API routes for HTMX requests
	api.GET("/inventory/low-stock-count", apiHandler.GetLowStockCount)
	api.GET("/sales/today", apiHandler.GetTodaySales)
	api.GET("/inventory/total-count", apiHandler.GetTotalInventoryCount)
	api.GET("/compatibility/vehicle-count", apiHandler.GetVehicleCount)
	api.GET("/inventory/low-stock", apiHandler.GetLowStockItems)
	api.GET("/sales/recent", apiHandler.GetRecentSales)
	api.GET("/sales/top-sellers", apiHandler.GetTopSellers)
	api.GET("/purchases/recent", apiHandler.GetRecentPurchases)
}
