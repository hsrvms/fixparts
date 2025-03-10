package repositories

import (
	"context"
	"errors"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/makes/models"
	vehicleModelModels "github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
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

func (r *PostgresVehicleMakeRepository) GetModelsByMake(ctx context.Context, makeID int) ([]*vehicleModelModels.VehicleModel, error) {
	query := `
		SELECT m.model_id, m.make_id, m.model_name, m.created_at, m.updated_at,
			   mk.make_name
		FROM vehicle_models m
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		WHERE m.make_id = $1
		ORDER BY m.model_name
	`

	rows, err := r.db.Pool.Query(ctx, query, makeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []*vehicleModelModels.VehicleModel
	for rows.Next() {
		model := &vehicleModelModels.VehicleModel{}
		err := rows.Scan(
			&model.ModelID,
			&model.MakeID,
			&model.ModelName,
			&model.CreatedAt,
			&model.UpdatedAt,
			&model.MakeName,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, rows.Err()
}
