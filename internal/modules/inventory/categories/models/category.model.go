package models

import (
	"time"
)

type Category struct {
	CategoryID       int       `json:"category_id" db:"category_id"`
	CategoryName     string    `json:"category_name" db:"category_name"`
	Description      *string   `json:"description,omitempty" db:"description"`
	ParentCategoryID *int      `json:"parent_category_id,omitempty" db:"parent_category_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`

	// Nested relationship (not from database)
	Subcategories []*Category `json:"subcategories,omitempty" db:"-"`
}

type CategoryTreeNode struct {
	Category      *Category           `json:"category"`
	Subcategories []*CategoryTreeNode `json:"subcategories,omitempty"`
}
