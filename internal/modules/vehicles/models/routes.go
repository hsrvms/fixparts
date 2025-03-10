package vehicleModels

import (
	vehicleMakeRepositories "github.com/hsrvms/fixparts/internal/modules/vehicles/makes/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/handlers"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresVehicleModelRepository(database)
	vehicleMakeRepo := vehicleMakeRepositories.NewPostgresVehicleMakeRepository(database)
	service := services.NewVehicleModelService(repo, vehicleMakeRepo)
	handler := handlers.NewVehicleModelHandler(service)

	models := api.Group("/models")
	models.GET("", handler.GetAllModels)
	models.GET("/:id", handler.GetModelByID)
	models.POST("", handler.CreateModel)
	models.PUT("/:id", handler.UpdateModel)
	models.DELETE("/:id", handler.DeleteModel)
	models.GET("/:modelId/submodels", handler.GetSubmodelsByModel) // Get submodels for a specific model
}
