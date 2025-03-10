package purchases

import (
	"github.com/hsrvms/fixparts/internal/modules/purchases/handlers"
	"github.com/hsrvms/fixparts/internal/modules/purchases/repositories"
	"github.com/hsrvms/fixparts/internal/modules/purchases/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresPurchaseRepository(database)
	service := services.NewPurchaseService(repo)
	handler := handlers.NewPurchaseHandler(service)

	purchases := api.Group("/purchases")
	purchases.GET("", handler.GetPurchases)
	purchases.GET("/:id", handler.GetPurchaseByID)
	purchases.POST("", handler.CreatePurchase)
	purchases.PUT("/:id", handler.UpdatePurchase)
	purchases.DELETE("/:id", handler.DeletePurchase)

	api.GET("/suppliers/:supplierId/purchases", handler.GetSupplierPurchases)
	api.GET("/items/:itemId/purchases", handler.GetItemPurchases)
}
