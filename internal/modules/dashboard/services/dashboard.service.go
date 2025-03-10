package services

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/dashboard/models"
	"github.com/hsrvms/fixparts/internal/modules/dashboard/repositories"
)

type dashboardService struct {
	repo repositories.DashboardRepository
}

func NewDashboardService(repo repositories.DashboardRepository) DashboardService {
	return &dashboardService{
		repo: repo,
	}
}

func (s *dashboardService) GetLowStockCount(ctx context.Context) (int, error) {
	return s.repo.GetLowStockCount(ctx)
}

func (s *dashboardService) GetTodaySales(ctx context.Context) (float64, error) {
	return s.repo.GetTodaySales(ctx)
}

func (s *dashboardService) GetTotalInventoryCount(ctx context.Context) (int, error) {
	return s.repo.GetTotalInventoryCount(ctx)
}

func (s *dashboardService) GetVehicleCount(ctx context.Context) (int, error) {
	return s.repo.GetVehicleCount(ctx)
}

func (s *dashboardService) GetLowStockItems(ctx context.Context) ([]*models.LowStockItem, error) {
	return s.repo.GetLowStockItems(ctx)
}

func (s *dashboardService) GetRecentSales(ctx context.Context) ([]*models.RecentSale, error) {
	return s.repo.GetRecentSales(ctx)
}

func (s *dashboardService) GetTopSellers(ctx context.Context) ([]*models.TopSeller, error) {
	return s.repo.GetTopSellers(ctx)
}

func (s *dashboardService) GetRecentPurchases(ctx context.Context) ([]*models.RecentPurchase, error) {
	return s.repo.GetRecentPurchases(ctx)
}
