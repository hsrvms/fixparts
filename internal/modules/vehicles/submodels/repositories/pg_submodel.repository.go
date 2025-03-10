package repositories

import (
	"context"
	"errors"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/jackc/pgx/v5"
)

type PostgresVehicleSubmodelRepository struct {
	db *db.Database
}

func NewPostgresVehicleSubmodelRepository(database *db.Database) VehicleSubmodelRepository {
	return &PostgresVehicleSubmodelRepository{
		db: database,
	}
}

func (r *PostgresVehicleSubmodelRepository) GetAllSubmodels(ctx context.Context) ([]*models.VehicleSubmodel, error) {
	query := `
		SELECT s.submodel_id, s.model_id, s.submodel_name, s.year_from, s.year_to,
			   s.engine_type, s.engine_displacement, s.fuel_type, s.transmission_type,
			   s.body_type, s.created_at, s.updated_at,
			   m.model_name, mk.make_name
		FROM vehicle_submodels s
		JOIN vehicle_models m ON s.model_id = m.model_id
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		ORDER BY mk.make_name, m.model_name, s.submodel_name
	`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submodels []*models.VehicleSubmodel
	for rows.Next() {
		submodel := &models.VehicleSubmodel{}
		err := rows.Scan(
			&submodel.SubmodelID,
			&submodel.ModelID,
			&submodel.SubmodelName,
			&submodel.YearFrom,
			&submodel.YearTo,
			&submodel.EngineType,
			&submodel.EngineDisplacement,
			&submodel.FuelType,
			&submodel.TransmissionType,
			&submodel.BodyType,
			&submodel.CreatedAt,
			&submodel.UpdatedAt,
			&submodel.ModelName,
			&submodel.MakeName,
		)
		if err != nil {
			return nil, err
		}
		submodels = append(submodels, submodel)
	}

	return submodels, rows.Err()
}

func (r *PostgresVehicleSubmodelRepository) GetSubmodelsByModel(ctx context.Context, modelID int) ([]*models.VehicleSubmodel, error) {
	query := `
		SELECT s.submodel_id, s.model_id, s.submodel_name, s.year_from, s.year_to,
			   s.engine_type, s.engine_displacement, s.fuel_type, s.transmission_type,
			   s.body_type, s.created_at, s.updated_at,
			   m.model_name, mk.make_name
		FROM vehicle_submodels s
		JOIN vehicle_models m ON s.model_id = m.model_id
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		WHERE s.model_id = $1
		ORDER BY s.year_from DESC, s.submodel_name
	`

	rows, err := r.db.Pool.Query(ctx, query, modelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var submodels []*models.VehicleSubmodel
	for rows.Next() {
		submodel := &models.VehicleSubmodel{}
		err := rows.Scan(
			&submodel.SubmodelID,
			&submodel.ModelID,
			&submodel.SubmodelName,
			&submodel.YearFrom,
			&submodel.YearTo,
			&submodel.EngineType,
			&submodel.EngineDisplacement,
			&submodel.FuelType,
			&submodel.TransmissionType,
			&submodel.BodyType,
			&submodel.CreatedAt,
			&submodel.UpdatedAt,
			&submodel.ModelName,
			&submodel.MakeName,
		)
		if err != nil {
			return nil, err
		}
		submodels = append(submodels, submodel)
	}

	return submodels, rows.Err()
}

func (r *PostgresVehicleSubmodelRepository) GetSubmodelByID(ctx context.Context, id int) (*models.VehicleSubmodel, error) {
	query := `
		SELECT s.submodel_id, s.model_id, s.submodel_name, s.year_from, s.year_to,
			   s.engine_type, s.engine_displacement, s.fuel_type, s.transmission_type,
			   s.body_type, s.created_at, s.updated_at,
			   m.model_name, mk.make_name
		FROM vehicle_submodels s
		JOIN vehicle_models m ON s.model_id = m.model_id
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		WHERE s.submodel_id = $1
	`

	submodel := &models.VehicleSubmodel{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&submodel.SubmodelID,
		&submodel.ModelID,
		&submodel.SubmodelName,
		&submodel.YearFrom,
		&submodel.YearTo,
		&submodel.EngineType,
		&submodel.EngineDisplacement,
		&submodel.FuelType,
		&submodel.TransmissionType,
		&submodel.BodyType,
		&submodel.CreatedAt,
		&submodel.UpdatedAt,
		&submodel.ModelName,
		&submodel.MakeName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return submodel, nil
}

func (r *PostgresVehicleSubmodelRepository) CreateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) (int, error) {
	query := `
		INSERT INTO vehicle_submodels (
			model_id, submodel_name, year_from, year_to,
			engine_type, engine_displacement, fuel_type,
			transmission_type, body_type
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING submodel_id
	`

	var id int
	err := r.db.Pool.QueryRow(
		ctx,
		query,
		submodel.ModelID,
		submodel.SubmodelName,
		submodel.YearFrom,
		submodel.YearTo,
		submodel.EngineType,
		submodel.EngineDisplacement,
		submodel.FuelType,
		submodel.TransmissionType,
		submodel.BodyType,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresVehicleSubmodelRepository) UpdateSubmodel(ctx context.Context, submodel *models.VehicleSubmodel) error {
	query := `
		UPDATE vehicle_submodels
		SET model_id = $2, submodel_name = $3, year_from = $4, year_to = $5,
			engine_type = $6, engine_displacement = $7, fuel_type = $8,
			transmission_type = $9, body_type = $10
		WHERE submodel_id = $1
	`

	result, err := r.db.Pool.Exec(
		ctx,
		query,
		submodel.SubmodelID,
		submodel.ModelID,
		submodel.SubmodelName,
		submodel.YearFrom,
		submodel.YearTo,
		submodel.EngineType,
		submodel.EngineDisplacement,
		submodel.FuelType,
		submodel.TransmissionType,
		submodel.BodyType,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("submodel not found")
	}

	return nil
}

func (r *PostgresVehicleSubmodelRepository) DeleteSubmodel(ctx context.Context, id int) error {
	query := `DELETE FROM vehicle_submodels WHERE submodel_id = $1`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("submodel not found")
	}

	return nil
}
