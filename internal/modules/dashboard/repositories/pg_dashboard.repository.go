package repositories

import (
	"context"

	"github.com/hsrvms/fixparts/internal/modules/dashboard/models"
	"github.com/hsrvms/fixparts/pkg/db"
)

type PostgresDashboardRepository struct {
	db *db.Database
}

func NewPostgresDashboardRepository(database *db.Database) DashboardRepository {
	return &PostgresDashboardRepository{
		db: database,
	}
}

func (r *PostgresDashboardRepository) GetLowStockCount(ctx context.Context) (int, error) {
	query := `
        SELECT COUNT(*)
        FROM items
        WHERE current_stock < minimum_stock
    `
	var count int
	err := r.db.Pool.QueryRow(ctx, query).Scan(&count)
	return count, err
}

func (r *PostgresDashboardRepository) GetTodaySales(ctx context.Context) (float64, error) {
	query := `
        SELECT COALESCE(SUM(total_price), 0)
        FROM sales
        WHERE DATE(created_at) = CURRENT_DATE
    `
	var total float64
	err := r.db.Pool.QueryRow(ctx, query).Scan(&total)
	return total, err
}

func (r *PostgresDashboardRepository) GetTotalInventoryCount(ctx context.Context) (int, error) {
	query := `
        SELECT COUNT(*)
        FROM items
        WHERE is_active = true
    `
	var count int
	err := r.db.Pool.QueryRow(ctx, query).Scan(&count)
	return count, err
}

func (r *PostgresDashboardRepository) GetVehicleCount(ctx context.Context) (int, error) {
	query := `
        SELECT COUNT(DISTINCT submodel_id)
        FROM compatibility
    `
	var count int
	err := r.db.Pool.QueryRow(ctx, query).Scan(&count)
	return count, err
}

func (r *PostgresDashboardRepository) GetLowStockItems(ctx context.Context) ([]*models.LowStockItem, error) {
	query := `
        SELECT part_number, item_name, current_stock, minimum_stock
        FROM items
        WHERE current_stock < minimum_stock
        ORDER BY current_stock ASC
        LIMIT 10
    `
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.LowStockItem
	for rows.Next() {
		item := &models.LowStockItem{}
		err := rows.Scan(
			&item.PartNumber,
			&item.Name,
			&item.Current,
			&item.Minimum,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresDashboardRepository) GetRecentSales(ctx context.Context) ([]*models.RecentSale, error) {
	query := `
		SELECT
            TO_CHAR(s.date, 'DD/MM/YYYY') as date,
            i.item_name as part,
            s.customer_name as customer,
            s.total_price as total
        FROM sales s
        JOIN items i ON s.item_id = i.item_id
        ORDER BY s.date DESC
        LIMIT 10
    `
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []*models.RecentSale
	for rows.Next() {
		sale := &models.RecentSale{}
		err := rows.Scan(
			&sale.Date,
			&sale.Part,
			&sale.Customer,
			&sale.Total,
		)
		if err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}
	return sales, rows.Err()
}

func (r *PostgresDashboardRepository) GetTopSellers(ctx context.Context) ([]*models.TopSeller, error) {
	query := `
		SELECT
            i.part_number,
            i.item_name as name,
            COUNT(*) as sold,
            SUM(s.total_price) as revenue
        FROM sales s
        JOIN items i ON s.item_id = i.item_id
        WHERE s.date >= CURRENT_DATE - INTERVAL '30 days'
        GROUP BY i.part_number, i.item_name
        ORDER BY sold DESC
        LIMIT 10
    `
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.TopSeller
	for rows.Next() {
		item := &models.TopSeller{}
		err := rows.Scan(
			&item.PartNumber,
			&item.Name,
			&item.Sold,
			&item.Revenue,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}

func (r *PostgresDashboardRepository) GetRecentPurchases(ctx context.Context) ([]*models.RecentPurchase, error) {
	query := `
		SELECT
            TO_CHAR(p.date, 'DD/MM/YYYY') as date,
            i.part_number,
            s.name as supplier,
            p.total_cost
        FROM purchases p
        JOIN items i ON p.item_id = i.item_id
        JOIN suppliers s ON p.supplier_id = s.supplier_id
        ORDER BY p.date DESC
        LIMIT 10
    `
	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []*models.RecentPurchase
	for rows.Next() {
		purchase := &models.RecentPurchase{}
		err := rows.Scan(
			&purchase.Date,
			&purchase.PartNumber,
			&purchase.Supplier,
			&purchase.Cost,
		)
		if err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}
	return purchases, rows.Err()
}
