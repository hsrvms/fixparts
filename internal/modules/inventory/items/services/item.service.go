package services

import (
	"context"
	"errors"
	"fmt"

	itemerrors "github.com/hsrvms/fixparts/internal/modules/inventory/items/errors"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/items/repositories"
)

type itemService struct {
	repo           repositories.ItemRepository
	barcodeService BarcodeService
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{
		repo:           repo,
		barcodeService: NewBarcodeService(),
	}
}

func (s *itemService) GetItems(ctx context.Context, filter *models.ItemFilter) ([]*models.Item, error) {
	return s.repo.GetItems(ctx, filter)
}

func (s *itemService) GetItemByID(ctx context.Context, id int) (*models.Item, error) {
	if id <= 0 {
		return nil, itemerrors.ErrInvalidItemID
	}

	item, err := s.repo.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, itemerrors.ErrItemNotFound
	}

	return item, nil
}

func (s *itemService) GetItemByPartNumber(ctx context.Context, partNumber string) (*models.Item, error) {
	if partNumber == "" {
		return nil, errors.New("part number is required")
	}

	return s.repo.GetItemByPartNumber(ctx, partNumber)
}

func (s *itemService) GetItemByBarcode(ctx context.Context, barcode string) (*models.Item, error) {
	if barcode == "" {
		return nil, errors.New("barcode is required")
	}

	return s.repo.GetItemByBarcode(ctx, barcode)
}

func (s *itemService) CreateItem(ctx context.Context, item *models.Item) (int, error) {
	// Validate basic item fields
	if err := s.validateItem(item); err != nil {
		return 0, err
	}

	// Generate barcode if not provided
	if item.Barcode == nil || *item.Barcode == "" {
		// Default to 0 if IDs are nil
		categoryID := 0
		if item.CategoryID != nil {
			categoryID = *item.CategoryID
		}
		supplierID := 0
		if item.SupplierID != nil {
			supplierID = *item.SupplierID
		}

		barcode, err := s.barcodeService.GenerateBarcode(categoryID, supplierID)
		if err != nil {
			return 0, fmt.Errorf("failed to generate barcode: %w", err)
		}
		item.Barcode = &barcode
	}

	// Check for duplicate part number
	existing, err := s.repo.GetItemByPartNumber(ctx, item.PartNumber)
	if err != nil {
		return 0, err
	}
	if existing != nil {
		return 0, itemerrors.ErrDuplicatePartNumber
	}

	// Check for duplicate barcode
	if item.Barcode != nil {
		existing, err = s.repo.GetItemByBarcode(ctx, *item.Barcode)
		if err != nil {
			return 0, err
		}
		if existing != nil {
			return 0, itemerrors.ErrDuplicateBarcode
		}
	}

	return s.repo.CreateItem(ctx, item)
}

func (s *itemService) UpdateItem(ctx context.Context, item *models.Item) error {
	if item.ItemID <= 0 {
		return itemerrors.ErrInvalidItemID
	}

	// Validate required fields
	if err := s.validateItem(item); err != nil {
		return err
	}

	// Check if item exists
	existing, err := s.repo.GetItemByID(ctx, item.ItemID)
	if err != nil {
		return err
	}
	if existing == nil {
		return itemerrors.ErrItemNotFound
	}

	// Check for duplicate part number if changed
	if item.PartNumber != existing.PartNumber {
		existingByPartNumber, err := s.repo.GetItemByPartNumber(ctx, item.PartNumber)
		if err != nil {
			return err
		}
		if existingByPartNumber != nil && existingByPartNumber.ItemID != item.ItemID {
			return itemerrors.ErrDuplicatePartNumber
		}
	}

	// Check for duplicate barcode if changed
	if item.Barcode != nil && *item.Barcode != "" &&
		(existing.Barcode == nil || *item.Barcode != *existing.Barcode) {
		existingByBarcode, err := s.repo.GetItemByBarcode(ctx, *item.Barcode)
		if err != nil {
			return err
		}
		if existingByBarcode != nil && existingByBarcode.ItemID != item.ItemID {
			return itemerrors.ErrDuplicateBarcode
		}
	}

	return s.repo.UpdateItem(ctx, item)
}

func (s *itemService) DeleteItem(ctx context.Context, id int) error {
	if id <= 0 {
		return itemerrors.ErrInvalidItemID
	}

	// Check if item exists
	existing, err := s.repo.GetItemByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return itemerrors.ErrItemNotFound
	}

	// You might want to add additional checks here:
	// - Check if item has any active sales/purchases
	// - Check if item is part of any active orders
	// - Instead of deleting, maybe just mark as inactive

	return s.repo.DeleteItem(ctx, id)
}

func (s *itemService) GetLowStockItems(ctx context.Context) ([]*models.Item, error) {
	return s.repo.GetLowStockItems(ctx)
}

func (s *itemService) validateItem(item *models.Item) error {
	if item.PartNumber == "" {
		return errors.New("part number is required")
	}
	if item.Description == "" {
		return errors.New("description is required")
	}
	if item.BuyPrice <= 0 {
		return errors.New("buy price must be greater than 0")
	}
	if item.SellPrice <= 0 {
		return errors.New("sell price must be greater than 0")
	}
	if item.CurrentStock < 0 {
		return errors.New("current stock cannot be negative")
	}
	if item.MinimumStock < 0 {
		return errors.New("minimum stock cannot be negative")
	}
	return nil
}
