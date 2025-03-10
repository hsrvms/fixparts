package saleErrors

import "errors"

var (
	ErrSaleNotFound               = errors.New("sale not found")
	ErrInvalidSaleID              = errors.New("invalid sale ID")
	ErrInvalidItemID              = errors.New("invalid item ID")
	ErrInvalidQuantity            = errors.New("quantity must be greater than 0")
	ErrInvalidPricePerUnit        = errors.New("price per unit must be greater than 0")
	ErrDuplicateTransactionNumber = errors.New("transaction number already exists")
	ErrInvalidDate                = errors.New("sale date cannot be in the future")
	ErrInsufficientStock          = errors.New("insufficient stock for sale")
	ErrInvalidCustomerEmail       = errors.New("invalid customer email format")
)
