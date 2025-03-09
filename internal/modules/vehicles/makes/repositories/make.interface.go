package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/models"
)

type VehicleMakeRepository interface {
	GetAllMakes(ctx context.Context) ([]*models.VehicleMake, error)
	GetMakeByID(ctx context.Context, id int) (*models.VehicleMake, error)
	CreateMake(ctx context.Context, make *models.VehicleMake) (int, error)
	UpdateMake(ctx context.Context, make *models.VehicleMake) error
	DeleteMake(ctx context.Context, id int) error
}
