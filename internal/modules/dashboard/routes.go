package dashboard

import (
	handlers "github.com/hsrvms/fixparts/internal/modules/dashboard/handlers"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	e.GET("/", handlers.ViewHandler)

	// // Initialize repository
	//    repo := repositories.NewPostgresDashboardRepository(database)

	//    // Initialize service
	//    service := services.NewDashboardService(repo)

	//    // Initialize handler
	//    handler := handlers.NewDashboardHandler(service)

	//    // Main dashboard page route
	//    e.GET("/", handler.RenderDashboard)

	//    // API routes for HTMX requests
	//    api.GET("/stats", handler.GetStats)
	//    api.GET("/stats/low-stock-count", handler.GetStats)
	//    api.GET("/stats/today-sales", handler.GetStats)
	//    api.GET("/stats/active-items", handler.GetStats)
	//    api.GET("/stats/supplier-count", handler.GetStats)
	//    api.GET("/activities/recent", handler.GetRecentActivities)
	//    api.GET("/inventory/low-stock", handler.GetLowStockItems)
}
