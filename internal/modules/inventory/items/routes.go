package items

import (
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/handlers"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/repositories"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresItemRepository(database)
	service := services.NewItemService(repo)
	handler := handlers.NewItemHandler(service)

	items := api.Group("/items")
	items.GET("", handler.GetItems)
	items.GET("/low-stock", handler.GetLowStockItems)
	items.GET("/:id", handler.GetItemByID)
	items.GET("/barcode/:barcode", handler.GetItemByBarcode)
	items.POST("", handler.CreateItem)
	items.PUT("/:id", handler.UpdateItem)
	items.DELETE("/:id", handler.DeleteItem)
	items.GET("/barcode/:barcode/image", handler.GetBarcodeImage)
}
