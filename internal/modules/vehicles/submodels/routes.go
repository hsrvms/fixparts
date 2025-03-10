package vehicleSubmodels

import (
	vehicleModelRepositories "github.com/hsrvms/fixparts/internal/modules/vehicles/models/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/handlers"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresVehicleSubmodelRepository(database)
	vehicleModelRepo := vehicleModelRepositories.NewPostgresVehicleModelRepository(database)
	service := services.NewVehicleSubmodelService(repo, vehicleModelRepo)
	handler := handlers.NewVehicleSubmodelHandler(service)

	submodels := api.Group("/submodels")
	submodels.GET("", handler.GetAllSubmodels)
	submodels.GET("/:id", handler.GetSubmodelByID)
	submodels.POST("", handler.CreateSubmodel)
	submodels.PUT("/:id", handler.UpdateSubmodel)
	submodels.DELETE("/:id", handler.DeleteSubmodel)
}
