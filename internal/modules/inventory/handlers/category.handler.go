package handlers

import (
	"net/http"
	"strconv"

	"github.com/hsrvms/fixparts/internal/modules/inventory/models"
	"github.com/hsrvms/fixparts/internal/modules/inventory/services"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetAllCategories(c echo.Context) error {
	ctx := c.Request().Context()
	categories, err := h.service.GetAllCategories(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategoryByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	ctx := c.Request().Context()
	category, err := h.service.GetCategoryByID(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if category == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Category not found")
	}

	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) GetSubcategories(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	ctx := c.Request().Context()
	subcategories, err := h.service.GetSubcategories(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subcategories)
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	id, err := h.service.CreateCategory(ctx, category)
	if err != nil {
		switch err {
		case services.ErrParentCategoryNotFound:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	category.CategoryID = id
	return c.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category.CategoryID = id

	ctx := c.Request().Context()
	err = h.service.UpdateCategory(ctx, category)
	if err != nil {
		switch err {
		case services.ErrCategoryNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case services.ErrParentCategoryNotFound, services.ErrCircularReference:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid category ID")
	}

	ctx := c.Request().Context()
	err = h.service.DeleteCategory(ctx, id)
	if err != nil {
		switch err {
		case services.ErrCategoryNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case services.ErrCategoryHasSubcategories:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CategoryHandler) GetCategoryTree(c echo.Context) error {
	ctx := c.Request().Context()
	tree, err := h.service.GetCategoryTree(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tree)
}
