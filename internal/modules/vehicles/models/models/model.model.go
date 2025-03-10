package models

import "time"

type VehicleModel struct {
	ModelID   int       `json:"model_id" db:"model_id"`
	MakeID    int       `json:"make_id" db:"make_id"`
	ModelName string    `json:"model_name" db:"model_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Additional fields for API responses
	MakeName string `json:"make_name,omitempty" db:"-"`
}
