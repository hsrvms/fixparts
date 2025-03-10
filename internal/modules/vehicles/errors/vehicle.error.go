package vehicleErrors

import "errors"

var (
	ErrMakeNotFound     = errors.New("make not found")
	ErrModelNotFound    = errors.New("model not found")
	ErrSubmodelNotFound = errors.New("submodel not found")
	ErrInvalidMakeID    = errors.New("invalid make ID")
	ErrInvalidModelID   = errors.New("invalid model ID")
)
