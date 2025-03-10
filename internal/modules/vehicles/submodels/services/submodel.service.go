package services

import (
	"context"
	"errors"

	vehicleErrors "github.com/hsrvms/fixparts/internal/modules/vehicles/errors"
	vehicleModelRepositories "github.com/hsrvms/fixparts/internal/modules/vehicles/models/repositories"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/repositories"
)

type vehicleSubmodelService struct {
	repo             repositories.VehicleSubmodelRepository
	vehicleModelRepo vehicleModelRepositories.VehicleModelRepository
}

func NewVehicleSubmodelService(
	repo repositories.VehicleSubmodelRepository,
	vehicleModelRepo vehicleModelRepositories.VehicleModelRepository,
) VehicleSubmodelService {
	return &vehicleSubmodelService{
		repo:             repo,
		vehicleModelRepo: vehicleModelRepo,
	}
}

func (s *vehicleSubmodelService) GetAllSubmodels(ctx context.Context) ([]*models.VehicleSubmodel, error) {
	return s.repo.GetAllSubmodels(ctx)
}

func (s *vehicleSubmodelService) GetSubmodelByID(ctx context.Context, id int) (*models.VehicleSubmodel, error) {
	if id <= 0 {
		return nil, errors.New("invalid submodel ID")
	}

	submodel, err := s.repo.GetSubmodelByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if submodel == nil {
		return nil, vehicleErrors.ErrSubmodelNotFound
	}

	return submodel, nil
}

func (s *vehicleSubmodelService) CreateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) (int, error) {
	// Validate required fields
	if err := s.validateSubmodel(submodel); err != nil {
		return 0, err
	}

	// Verify model exists
	model, err := s.vehicleModelRepo.GetModelByID(ctx, submodel.ModelID)
	if err != nil {
		return 0, err
	}
	if model == nil {
		return 0, vehicleErrors.ErrModelNotFound
	}

	// Validate year range
	if submodel.YearTo != nil && *submodel.YearTo < submodel.YearFrom {
		return 0, errors.New("end year cannot be earlier than start year")
	}

	return s.repo.CreateSubmodel(ctx, submodel)
}

func (s *vehicleSubmodelService) UpdateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) error {
	if submodel.SubmodelID <= 0 {
		return errors.New("invalid submodel ID")
	}

	// Validate required fields
	if err := s.validateSubmodel(submodel); err != nil {
		return err
	}

	// Check if submodel exists
	existing, err := s.repo.GetSubmodelByID(ctx, submodel.SubmodelID)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrSubmodelNotFound
	}

	// Verify model exists if model ID is being changed
	if existing.ModelID != submodel.ModelID {
		model, err := s.vehicleModelRepo.GetModelByID(ctx, submodel.ModelID)
		if err != nil {
			return err
		}
		if model == nil {
			return vehicleErrors.ErrModelNotFound
		}
	}

	// Validate year range
	if submodel.YearTo != nil && *submodel.YearTo < submodel.YearFrom {
		return errors.New("end year cannot be earlier than start year")
	}

	return s.repo.UpdateSubmodel(ctx, submodel)
}

func (s *vehicleSubmodelService) DeleteSubmodel(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid submodel ID")
	}

	// Check if submodel exists
	existing, err := s.repo.GetSubmodelByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return vehicleErrors.ErrSubmodelNotFound
	}

	// You might want to check for dependencies (e.g., parts compatibility) before deletion
	// This would require additional repository methods

	return s.repo.DeleteSubmodel(ctx, id)
}

// Helper function for submodel validation
func (s *vehicleSubmodelService) validateSubmodel(submodel *models.VehicleSubmodel) error {
	if submodel.SubmodelName == "" {
		return errors.New("submodel name is required")
	}
	if submodel.YearFrom <= 0 {
		return errors.New("valid start year is required")
	}
	if submodel.EngineType == "" {
		return errors.New("engine type is required")
	}
	if submodel.EngineDisplacement <= 0 {
		return errors.New("valid engine displacement is required")
	}
	if submodel.FuelType == "" {
		return errors.New("fuel type is required")
	}
	if submodel.TransmissionType == "" {
		return errors.New("transmission type is required")
	}
	if submodel.BodyType == "" {
		return errors.New("body type is required")
	}
	return nil
}
