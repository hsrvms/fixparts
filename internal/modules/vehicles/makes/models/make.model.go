package models

import "time"

type VehicleMake struct {
	MakeID    int       `json:"make_id" db:"make_id"`
	MakeName  string    `json:"make_name" db:"make_name"`
	Country   *string   `json:"country" db:"country"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
