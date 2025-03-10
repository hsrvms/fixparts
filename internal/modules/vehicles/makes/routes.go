package vehicleMakes

import (
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/handlers"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresVehicleMakeRepository(database)
	service := services.NewVehicleMakeService(repo)
	handler := handlers.NewVehicleMakeHandler(service)

	makes := api.Group("/makes")
	makes.GET("", handler.GetAllMakes)
	makes.GET("/:id", handler.GetMakeByID)
	makes.POST("", handler.CreateMake)
	makes.PUT("/:id", handler.UpdateMake)
	makes.DELETE("/:id", handler.DeleteMake)
	makes.GET("/:makeId/models", handler.GetModelsByMake) // Get models for a specific make

}
