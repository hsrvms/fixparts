package compatibility

import (
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/handlers"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/repositories"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresCompatibilityRepository(database)
	service := services.NewCompatibilityService(repo)
	handler := handlers.NewCompatibilityHandler(service)

	items := api.Group("/items")
	items.GET("/:itemId/compatibilities", handler.GetCompatibilities)
	items.POST("/:itemId/compatibilities", handler.AddCompatibility)
	items.DELETE("/:itemId/compatibilities/:submodelId", handler.RemoveCompatibility)
	api.GET("/submodels/:submodelId/compatible-items", handler.GetCompatibleItems)
}
