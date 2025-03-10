package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/suppliers/models"
)

type SupplierService interface {
	GetAll(ctx context.Context, filter *models.SupplierFilter) ([]*models.Supplier, error)
	GetByID(ctx context.Context, id int) (*models.Supplier, error)
	Create(ctx context.Context, supplier *models.Supplier) (int, error)
	Update(ctx context.Context, supplier *models.Supplier) error
	Delete(ctx context.Context, id int) error
}
