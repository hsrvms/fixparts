package services

import (
	"context"

	compatibilityerrors "github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/errors"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/repositories"
	itemmodels "github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
	itemRepositories "github.com/hsrvms/fixparts/internal/modules/inventory/items/repositories"
)

type compatibilityService struct {
	repo     repositories.CompatibilityRepository
	itemRepo itemRepositories.ItemRepository
}

func NewCompatibilityService(repo repositories.CompatibilityRepository) CompatibilityService {
	return &compatibilityService{
		repo: repo,
	}
}

func (s *compatibilityService) GetCompatibilities(ctx context.Context, itemID int) ([]*models.Compatibility, error) {
	if itemID <= 0 {
		return nil, compatibilityerrors.ErrInvalidItemID
	}

	return s.repo.GetCompatibilities(ctx, itemID)
}

func (s *compatibilityService) AddCompatibility(ctx context.Context, compatibility *models.Compatibility) (int, error) {
	if compatibility.ItemID <= 0 {
		return 0, compatibilityerrors.ErrInvalidItemID
	}
	if compatibility.SubmodelID <= 0 {
		return 0, compatibilityerrors.ErrInvalidSubmodelID
	}

	// Check if item exists
	item, err := s.itemRepo.GetItemByID(ctx, compatibility.ItemID)
	if err != nil {
		return 0, err
	}
	if item == nil {
		return 0, compatibilityerrors.ErrItemNotFound
	}

	// Check if compatibility already exists
	compatibilities, err := s.repo.GetCompatibilities(ctx, compatibility.ItemID)
	if err != nil {
		return 0, err
	}

	for _, existing := range compatibilities {
		if existing.SubmodelID == compatibility.SubmodelID {
			return 0, compatibilityerrors.ErrCompatibilityExists
		}
	}

	return s.repo.AddCompatibility(ctx, compatibility)
}

func (s *compatibilityService) RemoveCompatibility(ctx context.Context, itemID, submodelID int) error {
	if itemID <= 0 {
		return compatibilityerrors.ErrInvalidItemID
	}
	if submodelID <= 0 {
		return compatibilityerrors.ErrInvalidSubmodelID
	}

	return s.repo.RemoveCompatibility(ctx, itemID, submodelID)
}

func (s *compatibilityService) GetCompatibleItems(ctx context.Context, submodelID int) ([]*itemmodels.Item, error) {
	if submodelID <= 0 {
		return nil, compatibilityerrors.ErrInvalidSubmodelID
	}

	return s.repo.GetCompatibleItems(ctx, submodelID)
}
