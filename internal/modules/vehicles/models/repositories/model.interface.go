package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
	vehicleSubModels "github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
)

type VehicleModelRepository interface {
	GetAllModels(ctx context.Context) ([]*models.VehicleModel, error)
	GetModelByID(ctx context.Context, id int) (*models.VehicleModel, error)
	CreateModel(ctx context.Context, model *models.VehicleModel) (int, error)
	UpdateModel(ctx context.Context, model *models.VehicleModel) error
	DeleteModel(ctx context.Context, id int) error

	GetSubmodelsByModel(ctx context.Context, modelID int) ([]*vehicleSubModels.VehicleSubmodel, error)
}
