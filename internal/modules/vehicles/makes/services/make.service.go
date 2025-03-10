package services

import (
	"context"
	"errors"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/repositories"
	vehicleModelModels "github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
)

type vehicleMakeService struct {
	repo repositories.VehicleMakeRepository
}

func NewVehicleMakeService(repo repositories.VehicleMakeRepository) VehicleMakeService {
	return &vehicleMakeService{
		repo: repo,
	}
}

func (s *vehicleMakeService) GetAllMakes(ctx context.Context) ([]*models.VehicleMake, error) {
	return s.repo.GetAllMakes(ctx)
}

func (s *vehicleMakeService) GetMakeByID(ctx context.Context, id int) (*models.VehicleMake, error) {
	if id <= 0 {
		return nil, vehicleErrors.ErrInvalidMakeID
	}

	make, err := s.repo.GetMakeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if make == nil {
		return nil, vehicleErrors.ErrMakeNotFound
	}

	return make, nil
}

func (s *vehicleMakeService) CreateMake(ctx context.Context, make *models.VehicleMake) (int, error) {
	// Validate required fields
	if make.MakeName == "" {
		return 0, errors.New("make name is required")
	}

	return s.repo.CreateMake(ctx, make)
}

func (s *vehicleMakeService) UpdateMake(ctx context.Context, make *models.VehicleMake) error {
	if make.MakeID <= 0 {
		return vehicleErrors.ErrInvalidMakeID
	}

	if make.MakeName == "" {
		return errors.New("make name is required")
	}

	// Check if make exists
	existing, err := s.repo.GetMakeByID(ctx, make.MakeID)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrMakeNotFound
	}

	return s.repo.UpdateMake(ctx, make)
}

func (s *vehicleMakeService) DeleteMake(ctx context.Context, id int) error {
	if id <= 0 {
		return vehicleErrors.ErrInvalidMakeID
	}

	// Check if make exists
	existing, err := s.repo.GetMakeByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrMakeNotFound
	}

	// Check if make has any models
	models, err := s.repo.GetModelsByMake(ctx, id)
	if err != nil {
		return err
	}
	if len(models) > 0 {
		return errors.New("cannot delete make with existing models")
	}

	return s.repo.DeleteMake(ctx, id)
}

func (s *vehicleMakeService) GetModelsByMake(ctx context.Context, makeID int) ([]*vehicleModelModels.VehicleModel, error) {
	if makeID <= 0 {
		return nil, vehicleErrors.ErrInvalidMakeID
	}

	// Verify make exists
	make, err := s.repo.GetMakeByID(ctx, makeID)
	if err != nil {
		return nil, err
	}
	if make == nil {
		return nil, vehicleErrors.ErrMakeNotFound
	}

	return s.repo.GetModelsByMake(ctx, makeID)
}
