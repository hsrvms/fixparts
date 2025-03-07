package compatibilityerrors

import "errors"

var (
	ErrItemNotFound        = errors.New("item not found")
	ErrInvalidItemID       = errors.New("invalid item ID")
	ErrInvalidSubmodelID   = errors.New("invalid submodel ID")
	ErrCompatibilityExists = errors.New("compatibility already exists")
)
