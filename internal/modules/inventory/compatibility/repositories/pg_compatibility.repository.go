package repositories

import (
	"context"
	"errors"

	"github.com/hsrvms/fixparts/internal/modules/inventory/compatibility/models"
	itemmodels "github.com/hsrvms/fixparts/internal/modules/inventory/items/models"
	"github.com/hsrvms/fixparts/pkg/db"
)

type PostgresCompatibilityRepository struct {
	db *db.Database
}

func NewPostgresCompatibilityRepository(database *db.Database) CompatibilityRepository {
	return &PostgresCompatibilityRepository{
		db: database,
	}
}

func (r *PostgresCompatibilityRepository) GetCompatibilities(ctx context.Context, itemID int) ([]*models.Compatibility, error) {
	query := `
        SELECT
            c.compat_id, c.item_id, c.submodel_id, c.notes, c.created_at,
            m.model_name, mk.make_name, s.submodel_name
        FROM compatibility c
        JOIN vehicle_submodels s ON c.submodel_id = s.submodel_id
        JOIN vehicle_models m ON s.model_id = m.model_id
        JOIN vehicle_makes mk ON m.make_id = mk.make_id
        WHERE c.item_id = $1
        ORDER BY mk.make_name, m.model_name, s.submodel_name
    `

	rows, err := r.db.Pool.Query(ctx, query, itemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compatibilities []*models.Compatibility
	for rows.Next() {
		compatibility := &models.Compatibility{}
		err := rows.Scan(
			&compatibility.CompatID,
			&compatibility.ItemID,
			&compatibility.SubmodelID,
			&compatibility.Notes,
			&compatibility.CreatedAt,
			&compatibility.ModelName,
			&compatibility.MakeName,
			&compatibility.SubmodelName,
		)
		if err != nil {
			return nil, err
		}
		compatibilities = append(compatibilities, compatibility)
	}

	return compatibilities, rows.Err()
}

func (r *PostgresCompatibilityRepository) AddCompatibility(ctx context.Context, compatibility *models.Compatibility) (int, error) {
	query := `
        INSERT INTO compatibility (item_id, submodel_id, notes)
        VALUES ($1, $2, $3)
        RETURNING compat_id
    `

	var id int
	err := r.db.Pool.QueryRow(
		ctx, query,
		compatibility.ItemID,
		compatibility.SubmodelID,
		compatibility.Notes,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresCompatibilityRepository) RemoveCompatibility(ctx context.Context, itemID, submodelID int) error {
	query := `
        DELETE FROM compatibility
        WHERE item_id = $1 AND submodel_id = $2
    `

	result, err := r.db.Pool.Exec(ctx, query, itemID, submodelID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("compatibility not found")
	}

	return nil
}

func (r *PostgresCompatibilityRepository) GetCompatibleItems(ctx context.Context, submodelID int) ([]*itemmodels.Item, error) {
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
        JOIN compatibility comp ON i.item_id = comp.item_id
        WHERE comp.submodel_id = $1 AND i.is_active = true
        ORDER BY i.part_number
    `

	rows, err := r.db.Pool.Query(ctx, query, submodelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*itemmodels.Item
	for rows.Next() {
		item := &itemmodels.Item{}
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

func (r *PostgresCompatibilityRepository) GetLowStockItems(ctx context.Context) ([]*itemmodels.Item, error) {
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

	var items []*itemmodels.Item
	for rows.Next() {
		item := &itemmodels.Item{}
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
