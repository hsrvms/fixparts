package services

import (
	"context"

	categoryerrors "github.com/hsrvms/fixparts/internal/modules/inventory/categories/errors"
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/repositories"
)

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *categoryService) GetCategoryByID(ctx context.Context, id int) (*models.Category, error) {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category != nil {
		subcategories, err := s.repo.GetSubcategories(ctx, category.CategoryID)
		if err != nil {
			return nil, err
		}
		category.Subcategories = subcategories
	}

	return category, nil
}

func (s *categoryService) GetSubcategories(ctx context.Context, parentID int) ([]*models.Category, error) {
	return s.repo.GetSubcategories(ctx, parentID)
}

func (s *categoryService) CreateCategory(ctx context.Context, category *models.Category) (int, error) {
	// Add any business logic/validation here before calling repository

	// For example, we might want to check if the parent category exists
	if category.ParentCategoryID != nil {
		parent, err := s.repo.GetByID(ctx, *category.ParentCategoryID)
		if err != nil {
			return 0, err
		}
		if parent == nil {
			return 0, categoryerrors.ErrParentCategoryNotFound
		}
	}

	return s.repo.Create(ctx, category)
}

// UpdateCategory updates an existing category
func (s *categoryService) UpdateCategory(ctx context.Context, category *models.Category) error {
	// Check if category exists
	existingCategory, err := s.repo.GetByID(ctx, category.CategoryID)
	if err != nil {
		return err
	}
	if existingCategory == nil {
		return categoryerrors.ErrCategoryNotFound
	}

	// Check if parent category exists (if one is specified)
	if category.ParentCategoryID != nil {
		parent, err := s.repo.GetByID(ctx, *category.ParentCategoryID)
		if err != nil {
			return err
		}
		if parent == nil {
			return categoryerrors.ErrParentCategoryNotFound
		}

		// Prevent circular references
		if *category.ParentCategoryID == category.CategoryID {
			return categoryerrors.ErrCircularReference
		}
	}

	return s.repo.Update(ctx, category)
}

// DeleteCategory removes a category
func (s *categoryService) DeleteCategory(ctx context.Context, id int) error {
	// Check if category exists
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if category == nil {
		return categoryerrors.ErrCategoryNotFound
	}

	// Check if category has subcategories
	subcategories, err := s.repo.GetSubcategories(ctx, id)
	if err != nil {
		return err
	}
	if len(subcategories) > 0 {
		return categoryerrors.ErrCategoryHasSubcategories
	}

	return s.repo.Delete(ctx, id)
}

// GetCategoryTree returns a hierarchical structure of categories
func (s *categoryService) GetCategoryTree(ctx context.Context) ([]*models.CategoryTreeNode, error) {
	return s.repo.GetCategoryTree(ctx)
}
