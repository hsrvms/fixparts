package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/models"
)

type CategoryService interface {
	GetAllCategories(ctx context.Context) ([]*models.Category, error)
	GetCategoryByID(ctx context.Context, id int) (*models.Category, error)
	GetSubcategories(ctx context.Context, parentID int) ([]*models.Category, error)
	CreateCategory(ctx context.Context, category *models.Category) (int, error)
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id int) error
	GetCategoryTree(ctx context.Context) ([]*models.CategoryTreeNode, error)
}
