package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
)

type ItemService interface {
	GetItems(ctx context.Context, filter *models.ItemFilter) ([]*models.Item, error)
	GetItemByID(ctx context.Context, id int) (*models.Item, error)
	GetItemByPartNumber(ctx context.Context, partNumber string) (*models.Item, error)
	GetItemByBarcode(ctx context.Context, barcode string) (*models.Item, error)
	CreateItem(ctx context.Context, item *models.Item) (int, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id int) error
	GetLowStockItems(ctx context.Context) ([]*models.Item, error)
}
