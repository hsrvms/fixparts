package categories

import (
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/handlers"
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/repositories"
	"github.com/hsrvms/fixparts/internal/modules/inventory/categories/services"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, api *echo.Group, database *db.Database) {
	repo := repositories.NewPostgresCategoryRepository(database)
	service := services.NewCategoryService(repo)
	handler := handlers.NewCategoryHandler(service)

	categories := api.Group("/categories")
	categories.GET("", handler.GetAllCategories)
	categories.GET("/:id", handler.GetCategoryByID)
	categories.GET("/:id/subcategories", handler.GetSubcategories)
	categories.POST("", handler.CreateCategory)
	categories.PUT("/:id", handler.UpdateCategory)
	categories.DELETE("/:id", handler.DeleteCategory)
	categories.GET("/tree", handler.GetCategoryTree)
}
