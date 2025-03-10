package services

import (
	"context"
	"errors"
	"time"

	saleErrors "github.com/hsrvms/fixparts/internal/modules/sales/errors"
	"github.com/hsrvms/fixparts/internal/modules/sales/models"
	"github.com/hsrvms/fixparts/internal/modules/sales/repositories"
)

type saleService struct {
	repo repositories.SaleRepository
}

func NewSaleService(repo repositories.SaleRepository) SaleService {
	return &saleService{
		repo: repo,
	}
}

func (s *saleService) GetAll(ctx context.Context, filter *models.SaleFilter) ([]*models.Sale, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s *saleService) GetByID(ctx context.Context, id int) (*models.Sale, error) {
	if id <= 0 {
		return nil, saleErrors.ErrInvalidSaleID
	}

	sale, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if sale == nil {
		return nil, saleErrors.ErrSaleNotFound
	}

	return sale, nil
}

func (s *saleService) Create(ctx context.Context, sale *models.Sale) (int, error) {
	// Validate the sale
	if err := s.validateSale(sale); err != nil {
		return 0, err
	}

	// Check if transaction number is unique if provided
	if sale.TransactionNumber != "" {
		existing, err := s.repo.GetByTransactionNumber(ctx, sale.TransactionNumber)
		if err != nil {
			return 0, err
		}
		if existing != nil {
			return 0, saleErrors.ErrDuplicateTransactionNumber
		}
	}

	// Set date to current time if not provided
	if sale.Date.IsZero() {
		sale.Date = time.Now()
	}

	// Calculate total price if not provided
	if sale.TotalPrice == 0 {
		sale.TotalPrice = float64(sale.Quantity) * sale.PricePerUnit
	}

	return s.repo.Create(ctx, sale)
}

func (s *saleService) Update(ctx context.Context, sale *models.Sale) error {
	if sale.SaleID <= 0 {
		return saleErrors.ErrInvalidSaleID
	}

	// Validate the sale
	if err := s.validateSale(sale); err != nil {
		return err
	}

	// Check if sale exists
	existing, err := s.repo.GetByID(ctx, sale.SaleID)
	if err != nil {
		return err
	}
	if existing == nil {
		return saleErrors.ErrSaleNotFound
	}

	// Check if transaction number is unique if changed
	if sale.TransactionNumber != existing.TransactionNumber {
		existingByTxn, err := s.repo.GetByTransactionNumber(ctx, sale.TransactionNumber)
		if err != nil {
			return err
		}
		if existingByTxn != nil && existingByTxn.SaleID != sale.SaleID {
			return saleErrors.ErrDuplicateTransactionNumber
		}
	}

	// Recalculate total price
	sale.TotalPrice = float64(sale.Quantity) * sale.PricePerUnit

	return s.repo.Update(ctx, sale)
}

func (s *saleService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return saleErrors.ErrInvalidSaleID
	}

	// Check if sale exists
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return saleErrors.ErrSaleNotFound
	}

	return s.repo.Delete(ctx, id)
}

func (s *saleService) GetByTransactionNumber(ctx context.Context, transactionNumber string) (*models.Sale, error) {
	if transactionNumber == "" {
		return nil, errors.New("transaction number is required")
	}

	return s.repo.GetByTransactionNumber(ctx, transactionNumber)
}

func (s *saleService) GetItemSales(ctx context.Context, itemID int) ([]*models.Sale, error) {
	if itemID <= 0 {
		return nil, saleErrors.ErrInvalidItemID
	}

	return s.repo.GetItemSales(ctx, itemID)
}

func (s *saleService) GetCustomerSales(ctx context.Context, customerEmail string) ([]*models.Sale, error) {
	if customerEmail == "" {
		return nil, errors.New("customer email is required")
	}

	return s.repo.GetCustomerSales(ctx, customerEmail)
}

// Helper functions
func (s *saleService) validateSale(sale *models.Sale) error {
	if sale.ItemID <= 0 {
		return saleErrors.ErrInvalidItemID
	}
	if sale.Quantity <= 0 {
		return saleErrors.ErrInvalidQuantity
	}
	if sale.PricePerUnit <= 0 {
		return saleErrors.ErrInvalidPricePerUnit
	}
	if !sale.Date.IsZero() && sale.Date.After(time.Now()) {
		return saleErrors.ErrInvalidDate
	}

	// Additional validations could be added here:
	// - Check if item exists
	// - Check if item has sufficient stock
	// - Validate email format if provided
	// - etc.

	return nil
}
