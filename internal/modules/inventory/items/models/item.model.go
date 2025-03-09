package models

import "time"

type Item struct {
	ItemID         int       `json:"item_id" db:"item_id"`
	ItemName       string    `json:"item_name" db:"item_name"`
	PartNumber     string    `json:"part_number" db:"part_number"`
	Description    string    `json:"description" db:"description"`
	CategoryID     *int      `json:"category_id,omitempty" db:"category_id"`
	BuyPrice       float64   `json:"buy_price" db:"buy_price"`
	SellPrice      float64   `json:"sell_price" db:"sell_price"`
	CurrentStock   int       `json:"current_stock" db:"current_stock"`
	MinimumStock   int       `json:"minimum_stock" db:"minimum_stock"`
	Barcode        *string   `json:"barcode,omitempty" db:"barcode"`
	SupplierID     *int      `json:"supplier_id,omitempty" db:"supplier_id"`
	LocationAisle  *string   `json:"location_aisle,omitempty" db:"location_aisle"`
	LocationShelf  *string   `json:"location_shelf,omitempty" db:"location_shelf"`
	LocationBin    *string   `json:"location_bin,omitempty" db:"location_bin"`
	WeightKg       *float64  `json:"weight_kg,omitempty" db:"weight_kg"`
	DimensionsCm   *string   `json:"dimensions_cm,omitempty" db:"dimensions_cm"`
	WarrantyPeriod *string   `json:"warranty_period,omitempty" db:"warranty_period"`
	ImageURL       *string   `json:"image_url,omitempty" db:"image_url"`
	IsActive       bool      `json:"is_active" db:"is_active"`
	Notes          *string   `json:"notes,omitempty" db:"notes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`

	// Additional fields for API responses
	CategoryName *string `json:"category_name,omitempty" db:"-"`
	SupplierName *string `json:"supplier_name,omitempty" db:"-"`
}
