package sales

import (
	"github.com/hsrvms/fixparts/internal/modules/sales/handlers"
	"github.com/hsrvms/fixparts/internal/modules/sales/repositories"
	"github.com/hsrvms/fixparts/internal/modules/sales/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresSaleRepository(database)
	service := services.NewSaleService(repo)
	handler := handlers.NewSaleHandler(service)

	sales := api.Group("/sales")
	sales.GET("", handler.GetSales)
	sales.GET("/:id", handler.GetSaleByID)
	sales.POST("", handler.CreateSale)
	sales.PUT("/:id", handler.UpdateSale)
	sales.DELETE("/:id", handler.DeleteSale)
	sales.GET("/transaction/:transactionNumber", handler.GetByTransactionNumber)
	sales.GET("/customer/:customerEmail", handler.GetCustomerSales)
}
