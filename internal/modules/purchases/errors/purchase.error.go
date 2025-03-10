package purchaseErrors

import (
	"errors"
)

var (
	ErrPurchaseNotFound       = errors.New("purchase not found")
	ErrInvalidPurchaseID      = errors.New("invalid purchase ID")
	ErrInvalidSupplierID      = errors.New("invalid supplier ID")
	ErrInvalidItemID          = errors.New("invalid item ID")
	ErrInvalidQuantity        = errors.New("quantity must be greater than 0")
	ErrInvalidCostPerUnit     = errors.New("cost per unit must be greater than 0")
	ErrDuplicateInvoiceNumber = errors.New("invoice number already exists")
	ErrInvalidDate            = errors.New("purchase date cannot be in the future")
)
