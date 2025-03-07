package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/jackc/pgx/v5"
)

type PostgresItemRepository struct {
	db *db.Database
}

func NewPostgresItemRepository(database *db.Database) ItemRepository {
	return &PostgresItemRepository{
		db: database,
	}
}

func (r *PostgresItemRepository) GetItems(ctx context.Context, filter *models.ItemFilter) ([]*models.Item, error) {
	query := `
		SELECT
			i.item_id, i.part_number, i.description, i.category_id, i.buy_price,
			i.sell_price, i.current_stock, i.minimum_stock, i.barcode, i.supplier_id,
			i.location_aisle, i.location_shelf, i.location_bin, i.weight_kg,
			i.dimensions_cm, i.warranty_period, i.image_url, i.is_active, i.notes,
			i.created_at, i.updated_at,
			c.category_name, s.name as supplier_name
		FROM items i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN suppliers s ON i.supplier_id = s.supplier_id
		WHERE 1=1
	`

	params := []interface{}{}
	paramCount := 1

	// Build query based on filters
	if filter != nil {
		if filter.CategoryID != nil {
			query += fmt.Sprintf(" AND i.category_id = $%d", paramCount)
			params = append(params, *filter.CategoryID)
			paramCount++
		}

		if filter.SupplierID != nil {
			query += fmt.Sprintf(" AND i.supplier_id = $%d", paramCount)
			params = append(params, *filter.SupplierID)
			paramCount++
		}

		if filter.PartNumber != nil {
			query += fmt.Sprintf(" AND i.part_number ILIKE $%d", paramCount)
			params = append(params, "%"+*filter.PartNumber+"%")
			paramCount++
		}

		if filter.SearchTerm != nil {
			query += fmt.Sprintf(" AND (i.part_number ILIKE $%d OR i.description ILIKE $%d)", paramCount, paramCount)
			params = append(params, "%"+*filter.SearchTerm+"%")
			paramCount++
		}

		if filter.LowStock != nil && *filter.LowStock {
			query += " AND i.current_stock <= i.minimum_stock"
		}

		if filter.IsActive != nil {
			query += fmt.Sprintf(" AND i.is_active = $%d", paramCount)
			params = append(params, *filter.IsActive)
			paramCount++
		}
	}

	query += " ORDER BY i.part_number"

	rows, err := r.db.Pool.Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.ItemID, &item.PartNumber, &item.Description, &item.CategoryID,
			&item.BuyPrice, &item.SellPrice, &item.CurrentStock, &item.MinimumStock,
			&item.Barcode, &item.SupplierID, &item.LocationAisle, &item.LocationShelf,
			&item.LocationBin, &item.WeightKg, &item.DimensionsCm, &item.WarrantyPeriod,
			&item.ImageURL, &item.IsActive, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
			&item.CategoryName, &item.SupplierName,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *PostgresItemRepository) GetItemByID(ctx context.Context, id int) (*models.Item, error) {
	query := `
		SELECT
			i.item_id, i.part_number, i.description, i.category_id, i.buy_price,
			i.sell_price, i.current_stock, i.minimum_stock, i.barcode, i.supplier_id,
			i.location_aisle, i.location_shelf, i.location_bin, i.weight_kg,
			i.dimensions_cm, i.warranty_period, i.image_url, i.is_active, i.notes,
			i.created_at, i.updated_at,
			c.category_name, s.name as supplier_name
		FROM items i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN suppliers s ON i.supplier_id = s.supplier_id
		WHERE i.item_id = $1
	`

	item := &models.Item{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&item.ItemID, &item.PartNumber, &item.Description, &item.CategoryID,
		&item.BuyPrice, &item.SellPrice, &item.CurrentStock, &item.MinimumStock,
		&item.Barcode, &item.SupplierID, &item.LocationAisle, &item.LocationShelf,
		&item.LocationBin, &item.WeightKg, &item.DimensionsCm, &item.WarrantyPeriod,
		&item.ImageURL, &item.IsActive, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
		&item.CategoryName, &item.SupplierName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func (r *PostgresItemRepository) GetItemByPartNumber(ctx context.Context, partNumber string) (*models.Item, error) {
	query := `
		SELECT
			i.item_id, i.part_number, i.description, i.category_id, i.buy_price,
			i.sell_price, i.current_stock, i.minimum_stock, i.barcode, i.supplier_id,
			i.location_aisle, i.location_shelf, i.location_bin, i.weight_kg,
			i.dimensions_cm, i.warranty_period, i.image_url, i.is_active, i.notes,
			i.created_at, i.updated_at,
			c.category_name, s.name as supplier_name
		FROM items i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN suppliers s ON i.supplier_id = s.supplier_id
		WHERE i.part_number = $1
	`

	item := &models.Item{}
	err := r.db.Pool.QueryRow(ctx, query, partNumber).Scan(
		&item.ItemID, &item.PartNumber, &item.Description, &item.CategoryID,
		&item.BuyPrice, &item.SellPrice, &item.CurrentStock, &item.MinimumStock,
		&item.Barcode, &item.SupplierID, &item.LocationAisle, &item.LocationShelf,
		&item.LocationBin, &item.WeightKg, &item.DimensionsCm, &item.WarrantyPeriod,
		&item.ImageURL, &item.IsActive, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
		&item.CategoryName, &item.SupplierName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func (r *PostgresItemRepository) GetItemByBarcode(ctx context.Context, barcode string) (*models.Item, error) {
	query := `
		SELECT
			i.item_id, i.part_number, i.description, i.category_id, i.buy_price,
			i.sell_price, i.current_stock, i.minimum_stock, i.barcode, i.supplier_id,
			i.location_aisle, i.location_shelf, i.location_bin, i.weight_kg,
			i.dimensions_cm, i.warranty_period, i.image_url, i.is_active, i.notes,
			i.created_at, i.updated_at,
			c.category_name, s.name as supplier_name
		FROM items i
		LEFT JOIN categories c ON i.category_id = c.category_id
		LEFT JOIN suppliers s ON i.supplier_id = s.supplier_id
		WHERE i.barcode = $1
	`

	item := &models.Item{}
	err := r.db.Pool.QueryRow(ctx, query, barcode).Scan(
		&item.ItemID, &item.PartNumber, &item.Description, &item.CategoryID,
		&item.BuyPrice, &item.SellPrice, &item.CurrentStock, &item.MinimumStock,
		&item.Barcode, &item.SupplierID, &item.LocationAisle, &item.LocationShelf,
		&item.LocationBin, &item.WeightKg, &item.DimensionsCm, &item.WarrantyPeriod,
		&item.ImageURL, &item.IsActive, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
		&item.CategoryName, &item.SupplierName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return item, nil
}

func (r *PostgresItemRepository) CreateItem(ctx context.Context, item *models.Item) (int, error) {
	query := `
		INSERT INTO items (
			part_number, item_name, description, category_id, buy_price, sell_price,
			current_stock, minimum_stock, barcode, supplier_id, location_aisle,
			location_shelf, location_bin, weight_kg, dimensions_cm,
			warranty_period, image_url, is_active, notes
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18, $19
		)
		RETURNING item_id
	`

	var id int
	err := r.db.Pool.QueryRow(
		ctx, query,
		item.PartNumber, item.ItemName, item.Description, item.CategoryID, item.BuyPrice,
		item.SellPrice, item.CurrentStock, item.MinimumStock, item.Barcode,
		item.SupplierID, item.LocationAisle, item.LocationShelf, item.LocationBin,
		item.WeightKg, item.DimensionsCm, item.WarrantyPeriod, item.ImageURL,
		item.IsActive, item.Notes,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresItemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	query := `
		UPDATE items SET
			part_number = $2, description = $3, category_id = $4,
			buy_price = $5, sell_price = $6, current_stock = $7,
			minimum_stock = $8, barcode = $9, supplier_id = $10,
			location_aisle = $11, location_shelf = $12, location_bin = $13,
			weight_kg = $14, dimensions_cm = $15, warranty_period = $16,
			image_url = $17, is_active = $18, notes = $19
		WHERE item_id = $1
	`

	result, err := r.db.Pool.Exec(
		ctx, query,
		item.ItemID, item.PartNumber, item.Description, item.CategoryID,
		item.BuyPrice, item.SellPrice, item.CurrentStock, item.MinimumStock,
		item.Barcode, item.SupplierID, item.LocationAisle, item.LocationShelf,
		item.LocationBin, item.WeightKg, item.DimensionsCm, item.WarrantyPeriod,
		item.ImageURL, item.IsActive, item.Notes,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("item not found")
	}

	return nil
}

func (r *PostgresItemRepository) DeleteItem(ctx context.Context, id int) error {
	query := `DELETE FROM items WHERE item_id = $1`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("item not found")
	}

	return nil
}

func (r *PostgresItemRepository) GetLowStockItems(ctx context.Context) ([]*models.Item, error) {
	query := `
        SELECT
            i.item_id, i.part_number, i.description, i.category_id, i.buy_price,
            i.sell_price, i.current_stock, i.minimum_stock, i.barcode, i.supplier_id,
            i.location_aisle, i.location_shelf, i.location_bin, i.weight_kg,
            i.dimensions_cm, i.warranty_period, i.image_url, i.is_active, i.notes,
            i.created_at, i.updated_at,
            c.category_name, s.name as supplier_name
        FROM items i
        LEFT JOIN categories c ON i.category_id = c.category_id
        LEFT JOIN suppliers s ON i.supplier_id = s.supplier_id
        WHERE i.current_stock <= i.minimum_stock AND i.is_active = true
        ORDER BY i.current_stock ASC, i.part_number
    `

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*models.Item
	for rows.Next() {
		item := &models.Item{}
		err := rows.Scan(
			&item.ItemID, &item.PartNumber, &item.Description, &item.CategoryID,
			&item.BuyPrice, &item.SellPrice, &item.CurrentStock, &item.MinimumStock,
			&item.Barcode, &item.SupplierID, &item.LocationAisle, &item.LocationShelf,
			&item.LocationBin, &item.WeightKg, &item.DimensionsCm, &item.WarrantyPeriod,
			&item.ImageURL, &item.IsActive, &item.Notes, &item.CreatedAt, &item.UpdatedAt,
			&item.CategoryName, &item.SupplierName,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
