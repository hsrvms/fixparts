package inventory

import (
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {

	inventoryGroup := api.Group("/inventory")

	categories.RegisterRoutes(e, inventoryGroup, database)
	items.RegisterRoutes(e, inventoryGroup, database)
	compatibility.RegisterRoutes(e, inventoryGroup, database)

}
