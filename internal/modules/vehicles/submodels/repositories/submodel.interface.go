package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
)

type VehicleSubmodelRepository interface {
	GetAllSubmodels(ctx context.Context) ([]*models.VehicleSubmodel, error)
	GetSubmodelByID(ctx context.Context, id int) (*models.VehicleSubmodel, error)
	CreateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) (int, error)
	UpdateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) error
	DeleteSubmodel(ctx context.Context, id int) error
}
