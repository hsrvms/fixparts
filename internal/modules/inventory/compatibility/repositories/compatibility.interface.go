package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/models"
	itemmodels "github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
)

type CompatibilityRepository interface {
	GetCompatibilities(ctx context.Context, itemID int) ([]*models.Compatibility, error)
	AddCompatibility(ctx context.Context, compatibility *models.Compatibility) (int, error)
	RemoveCompatibility(ctx context.Context, itemID, submodelID int) error
	GetCompatibleItems(ctx context.Context, submodelID int) ([]*itemmodels.Item, error)
}
