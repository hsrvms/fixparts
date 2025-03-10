package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/dashboard/models"
)

type DashboardService interface {
	GetLowStockCount(ctx context.Context) (int, error)
	GetTodaySales(ctx context.Context) (float64, error)
	GetTotalInventoryCount(ctx context.Context) (int, error)
	GetVehicleCount(ctx context.Context) (int, error)
	GetLowStockItems(ctx context.Context) ([]*models.LowStockItem, error)
	GetRecentSales(ctx context.Context) ([]*models.RecentSale, error)
	GetTopSellers(ctx context.Context) ([]*models.TopSeller, error)
	GetRecentPurchases(ctx context.Context) ([]*models.RecentPurchase, error)
}
