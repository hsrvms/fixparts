package suppliers

import (
	"github.com/hsrvms/fixparts/internal/modules/suppliers/handlers"
	"github.com/hsrvms/fixparts/internal/modules/suppliers/repositories"
	"github.com/hsrvms/fixparts/internal/modules/suppliers/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresSupplierRepository(database)
	service := services.NewSupplierService(repo)
	handler := handlers.NewSupplierHandler(service)

	suppliers := api.Group("/suppliers")
	suppliers.GET("", handler.GetSuppliers)
	suppliers.GET("/:id", handler.GetSupplierByID)
	suppliers.POST("", handler.CreateSupplier)
	suppliers.PUT("/:id", handler.UpdateSupplier)
	suppliers.DELETE("/:id", handler.DeleteSupplier)
}
