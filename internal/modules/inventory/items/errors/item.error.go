package itemerrors

import "errors"

var (
	ErrItemNotFound        = errors.New("item not found")
	ErrDuplicatePartNumber = errors.New("part number already exists")
	ErrDuplicateBarcode    = errors.New("barcode already exists")
	ErrInvalidItemID       = errors.New("invalid item ID")
	ErrInvalidSubmodelID   = errors.New("invalid submodel ID")
	ErrCompatibilityExists = errors.New("compatibility already exists")
	ErrInvalidPrice        = errors.New("price must be greater than 0")
	ErrInvalidStock        = errors.New("stock cannot be negative")
)
