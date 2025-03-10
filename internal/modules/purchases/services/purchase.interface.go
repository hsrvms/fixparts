package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/purchases/models"
)

type PurchaseService interface {
	GetAll(ctx context.Context, filter *models.PurchaseFilter) ([]*models.Purchase, error)
	GetByID(ctx context.Context, id int) (*models.Purchase, error)
	Create(ctx context.Context, purchase *models.Purchase) (int, error)
	Update(ctx context.Context, purchase *models.Purchase) error
	Delete(ctx context.Context, id int) error
	GetSupplierPurchases(ctx context.Context, supplierID int) ([]*models.Purchase, error)
	GetItemPurchases(ctx context.Context, itemID int) ([]*models.Purchase, error)
}
