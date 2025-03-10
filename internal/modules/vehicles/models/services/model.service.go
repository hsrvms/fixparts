package services

import (
	"context"
	"errors"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
	vehicleMakeRepositories "github.com/hsrvms/fixparts/internal/modules/vehicles/makes/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/repositories"
	vehicleSubmodels "github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
)

type vehicleModelService struct {
	repo            repositories.VehicleModelRepository
	vehicleMakeRepo vehicleMakeRepositories.VehicleMakeRepository
}

func NewVehicleModelService(
	repo repositories.VehicleModelRepository,
	vehicleMakeRepo vehicleMakeRepositories.VehicleMakeRepository,
) VehicleModelService {
	return &vehicleModelService{
		repo:            repo,
		vehicleMakeRepo: vehicleMakeRepo,
	}
}

func (s *vehicleModelService) GetAllModels(ctx context.Context) ([]*models.VehicleModel, error) {
	return s.repo.GetAllModels(ctx)
}

func (s *vehicleModelService) GetModelByID(ctx context.Context, id int) (*models.VehicleModel, error) {
	if id <= 0 {
		return nil, vehicleErrors.ErrInvalidModelID
	}

	model, err := s.repo.GetModelByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, vehicleErrors.ErrModelNotFound
	}

	return model, nil
}

func (s *vehicleModelService) CreateModel(ctx context.Context, model *models.VehicleModel) (int, error) {
	// Validate required fields
	if model.ModelName == "" {
		return 0, errors.New("model name is required")
	}

	// Verify make exists
	make, err := s.vehicleMakeRepo.GetMakeByID(ctx, model.MakeID)
	if err != nil {
		return 0, err
	}
	if make == nil {
		return 0, vehicleErrors.ErrMakeNotFound
	}

	return s.repo.CreateModel(ctx, model)
}

func (s *vehicleModelService) UpdateModel(ctx context.Context, model *models.VehicleModel) error {
	if model.ModelID <= 0 {
		return vehicleErrors.ErrInvalidModelID
	}

	if model.ModelName == "" {
		return errors.New("model name is required")
	}

	// Check if model exists
	existing, err := s.repo.GetModelByID(ctx, model.ModelID)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrModelNotFound
	}

	// Verify make exists if make ID is being changed
	if existing.MakeID != model.MakeID {
		make, err := s.vehicleMakeRepo.GetMakeByID(ctx, model.MakeID)
		if err != nil {
			return err
		}
		if make == nil {
			return vehicleErrors.ErrMakeNotFound
		}
	}

	return s.repo.UpdateModel(ctx, model)
}

func (s *vehicleModelService) DeleteModel(ctx context.Context, id int) error {
	if id <= 0 {
		return vehicleErrors.ErrInvalidModelID
	}

	// Check if model exists
	existing, err := s.repo.GetModelByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrModelNotFound
	}

	// Check if model has any submodels
	submodels, err := s.repo.GetSubmodelsByModel(ctx, id)
	if err != nil {
		return err
	}
	if len(submodels) > 0 {
		return errors.New("cannot delete model with existing submodels")
	}

	return s.repo.DeleteModel(ctx, id)
}

func (s *vehicleModelService) GetSubmodelsByModel(ctx context.Context, modelID int) ([]*vehicleSubmodels.VehicleSubmodel, error) {
	if modelID <= 0 {
		return nil, vehicleErrors.ErrInvalidModelID
	}

	// Verify model exists
	model, err := s.repo.GetModelByID(ctx, modelID)
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, vehicleErrors.ErrModelNotFound
	}

	return s.repo.GetSubmodelsByModel(ctx, modelID)
}
