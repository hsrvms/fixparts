package repositories

import (
	"context"
	"errors"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/models"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/jackc/pgx/v5"
)

type PostgresVehicleMakeRepository struct {
	db *db.Database
}

func NewPostgresVehicleMakeRepository(database *db.Database) VehicleMakeRepository {
	return &PostgresVehicleMakeRepository{
		db: database,
	}
}

func (r *PostgresVehicleMakeRepository) GetAllMakes(ctx context.Context) ([]*models.VehicleMake, error) {
	query := `
		SELECT make_id, make_name, country, created_at, updated_at
		FROM vehicle_makes
		ORDER BY make_name
	`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var makes []*models.VehicleMake
	for rows.Next() {
		make := &models.VehicleMake{}
		err := rows.Scan(
			&make.MakeID,
			&make.MakeName,
			&make.Country,
			&make.CreatedAt,
			&make.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		makes = append(makes, make)
	}

	return makes, rows.Err()
}

func (r *PostgresVehicleMakeRepository) GetMakeByID(ctx context.Context, id int) (*models.VehicleMake, error) {
	query := `
		SELECT make_id, make_name, country, created_at, updated_at
		FROM vehicle_makes
		WHERE make_id = $1
	`

	make := &models.VehicleMake{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&make.MakeID,
		&make.MakeName,
		&make.Country,
		&make.CreatedAt,
		&make.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return make, nil
}

func (r *PostgresVehicleMakeRepository) CreateMake(ctx context.Context, make *models.VehicleMake) (int, error) {
	query := `
		INSERT INTO vehicle_makes (make_name, country)
		VALUES ($1, $2)
		RETURNING make_id
	`

	var id int
	err := r.db.Pool.QueryRow(ctx, query, make.MakeName, make.Country).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresVehicleMakeRepository) UpdateMake(ctx context.Context, make *models.VehicleMake) error {
	query := `
		UPDATE vehicle_makes
		SET make_name = $2, country = $3
		WHERE make_id = $1
	`

	result, err := r.db.Pool.Exec(ctx, query, make.MakeID, make.MakeName, make.Country)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("make not found")
	}

	return nil
}

func (r *PostgresVehicleMakeRepository) DeleteMake(ctx context.Context, id int) error {
	query := `DELETE FROM vehicle_makes WHERE make_id = $1`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("make not found")
	}

	return nil
}
