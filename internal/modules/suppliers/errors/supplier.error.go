package supplierErrors

import "errors"

var (
	ErrSupplierNotFound      = errors.New("supplier not found")
	ErrInvalidSupplierID     = errors.New("invalid supplier ID")
	ErrDuplicateSupplierName = errors.New("supplier name already exists")
	ErrSupplierHasItems      = errors.New("cannot delete supplier with associated items")
)
