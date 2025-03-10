package models

import "time"

type VehicleSubmodel struct {
	SubmodelID         int       `json:"submodel_id" db:"submodel_id"`
	ModelID            int       `json:"model_id" db:"model_id"`
	SubmodelName       string    `json:"submodel_name" db:"submodel_name"`
	YearFrom           int       `json:"year_from" db:"year_from"`
	YearTo             *int      `json:"year_to,omitempty" db:"year_to"`
	EngineType         string    `json:"engine_type" db:"engine_type"`
	EngineDisplacement float64   `json:"engine_displacement" db:"engine_displacement"`
	FuelType           string    `json:"fuel_type" db:"fuel_type"`
	TransmissionType   string    `json:"transmission_type" db:"transmission_type"`
	BodyType           string    `json:"body_type" db:"body_type"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`

	// Additional fields for API responses
	ModelName string `json:"model_name,omitempty" db:"-"`
	MakeName  string `json:"make_name,omitempty" db:"-"`
}
