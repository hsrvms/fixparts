package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/sales/models"
)

type SaleRepository interface {
	GetAll(ctx context.Context, filter *models.SaleFilter) ([]*models.Sale, error)
	GetByID(ctx context.Context, id int) (*models.Sale, error)
	Create(ctx context.Context, sale *models.Sale) (int, error)
	Update(ctx context.Context, sale *models.Sale) error
	Delete(ctx context.Context, id int) error
	GetByTransactionNumber(ctx context.Context, transactionNumber string) (*models.Sale, error)
	GetItemSales(ctx context.Context, itemID int) ([]*models.Sale, error)
	GetCustomerSales(ctx context.Context, customerEmail string) ([]*models.Sale, error)
}
