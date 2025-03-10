package vehicles

import (
	vehicleMakes "github.com/hsrvms/fixparts/internal/modules/vehicles/makes"
	vehicleModels "github.com/hsrvms/fixparts/internal/modules/vehicles/models"
	vehicleSubmodels "github.com/hsrvms/fixparts/internal/modules/vehicles/submodels"

	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {

	vehicleGroup := api.Group("/vehicles")

	vehicleMakes.RegisterRoutes(e, vehicleGroup, database)
	vehicleModels.RegisterRoutes(e, vehicleGroup, database)
	vehicleSubmodels.RegisterRoutes(e, vehicleGroup, database)

}
