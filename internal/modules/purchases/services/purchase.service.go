package services

import (
	"context"
	"time"

	purchaseErrors "github.com/hsrvms/fixparts/internal/modules/purchases/errors"
	"github.com/hsrvms/fixparts/internal/modules/purchases/models"
	"github.com/hsrvms/fixparts/internal/modules/purchases/repositories"
)

type purchaseService struct {
	repo repositories.PurchaseRepository
}

func NewPurchaseService(repo repositories.PurchaseRepository) PurchaseService {
	return &purchaseService{
		repo: repo,
	}
}

func (s *purchaseService) GetAll(ctx context.Context, filter *models.PurchaseFilter) ([]*models.Purchase, error) {
	return s.repo.GetAll(ctx, filter)
}

func (s *purchaseService) GetByID(ctx context.Context, id int) (*models.Purchase, error) {
	if id <= 0 {
		return nil, purchaseErrors.ErrInvalidPurchaseID
	}

	purchase, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if purchase == nil {
		return nil, purchaseErrors.ErrPurchaseNotFound
	}

	return purchase, nil
}

func (s *purchaseService) Create(ctx context.Context, purchase *models.Purchase) (int, error) {
	// Validate the purchase
	if err := s.validatePurchase(purchase); err != nil {
		return 0, err
	}

	// Check if invoice number is unique if provided
	if purchase.InvoiceNumber != nil && *purchase.InvoiceNumber != "" {
		existing, err := s.repo.GetByInvoiceNumber(ctx, *purchase.InvoiceNumber)
		if err != nil {
			return 0, err
		}
		if existing != nil {
			return 0, purchaseErrors.ErrDuplicateInvoiceNumber
		}
	}

	// Set date to current time if not provided
	if purchase.Date.IsZero() {
		purchase.Date = time.Now()
	}

	// Calculate total cost if not provided
	if purchase.TotalCost == 0 {
		purchase.TotalCost = float64(purchase.Quantity) * purchase.CostPerUnit
	}

	return s.repo.Create(ctx, purchase)
}

func (s *purchaseService) Update(ctx context.Context, purchase *models.Purchase) error {
	if purchase.PurchaseID <= 0 {
		return purchaseErrors.ErrInvalidPurchaseID
	}

	// Validate the purchase
	if err := s.validatePurchase(purchase); err != nil {
		return err
	}

	// Check if purchase exists
	existing, err := s.repo.GetByID(ctx, purchase.PurchaseID)
	if err != nil {
		return err
	}
	if existing == nil {
		return purchaseErrors.ErrPurchaseNotFound
	}

	// Check if invoice number is unique if changed
	if purchase.InvoiceNumber != nil && *purchase.InvoiceNumber != "" &&
		(existing.InvoiceNumber == nil || *purchase.InvoiceNumber != *existing.InvoiceNumber) {
		existingWithInvoice, err := s.repo.GetByInvoiceNumber(ctx, *purchase.InvoiceNumber)
		if err != nil {
			return err
		}
		if existingWithInvoice != nil && existingWithInvoice.PurchaseID != purchase.PurchaseID {
			return purchaseErrors.ErrDuplicateInvoiceNumber
		}
	}

	// Recalculate total cost
	purchase.TotalCost = float64(purchase.Quantity) * purchase.CostPerUnit

	return s.repo.Update(ctx, purchase)
}

func (s *purchaseService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return purchaseErrors.ErrInvalidPurchaseID
	}

	// Check if purchase exists
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return purchaseErrors.ErrPurchaseNotFound
	}

	return s.repo.Delete(ctx, id)
}

func (s *purchaseService) GetSupplierPurchases(ctx context.Context, supplierID int) ([]*models.Purchase, error) {
	if supplierID <= 0 {
		return nil, purchaseErrors.ErrInvalidSupplierID
	}
	return s.repo.GetSupplierPurchases(ctx, supplierID)
}

func (s *purchaseService) GetItemPurchases(ctx context.Context, itemID int) ([]*models.Purchase, error) {
	if itemID <= 0 {
		return nil, purchaseErrors.ErrInvalidItemID
	}
	return s.repo.GetItemPurchases(ctx, itemID)
}

// Helper functions
func (s *purchaseService) validatePurchase(purchase *models.Purchase) error {
	if purchase.SupplierID <= 0 {
		return purchaseErrors.ErrInvalidSupplierID
	}
	if purchase.ItemID <= 0 {
		return purchaseErrors.ErrInvalidItemID
	}
	if purchase.Quantity <= 0 {
		return purchaseErrors.ErrInvalidQuantity
	}
	if purchase.CostPerUnit <= 0 {
		return purchaseErrors.ErrInvalidCostPerUnit
	}
	if !purchase.Date.IsZero() && purchase.Date.After(time.Now()) {
		return purchaseErrors.ErrInvalidDate
	}
	return nil
}
