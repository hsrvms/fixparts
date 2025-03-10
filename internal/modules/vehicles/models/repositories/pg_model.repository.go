package repositories

import (
	"context"
	"errors"

	"github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
	vehicleModelModels "github.com/hsrvms/fixparts/internal/modules/vehicles/models/models"
	vehicleSubmodels "github.com/hsrvms/fixparts/internal/modules/vehicles/submodels/models"
	"github.com/hsrvms/fixparts/pkg/db"
	"github.com/jackc/pgx/v5"
)

type PostgresVehicleModelRepository struct {
	db *db.Database
}

func NewPostgresVehicleModelRepository(database *db.Database) VehicleModelRepository {
	return &PostgresVehicleModelRepository{
		db: database,
	}
}

func (r *PostgresVehicleModelRepository) GetAllModels(ctx context.Context) ([]*vehicleModelModels.VehicleModel, error) {
	query := `
		SELECT m.model_id, m.make_id, m.model_name, m.created_at, m.updated_at,
			   mk.make_name
		FROM vehicle_models m
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		ORDER BY mk.make_name, m.model_name
	`

	rows, err := r.db.Pool.Query(ctx, query)
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

func (r *PostgresVehicleModelRepository) GetModelByID(ctx context.Context, id int) (*vehicleModelModels.VehicleModel, error) {
	query := `
		SELECT m.model_id, m.make_id, m.model_name, m.created_at, m.updated_at,
			   mk.make_name
		FROM vehicle_models m
		JOIN vehicle_makes mk ON m.make_id = mk.make_id
		WHERE m.model_id = $1
	`

	model := &models.VehicleModel{}
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&model.ModelID,
		&model.MakeID,
		&model.ModelName,
		&model.CreatedAt,
		&model.UpdatedAt,
		&model.MakeName,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return model, nil
}

func (r *PostgresVehicleModelRepository) CreateModel(ctx context.Context, model *vehicleModelModels.VehicleModel) (int, error) {
	query := `
		INSERT INTO vehicle_models (make_id, model_name)
		VALUES ($1, $2)
		RETURNING model_id
	`

	var id int
	err := r.db.Pool.QueryRow(ctx, query, model.MakeID, model.ModelName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostgresVehicleModelRepository) UpdateModel(ctx context.Context, model *vehicleModelModels.VehicleModel) error {
	query := `
		UPDATE vehicle_models
		SET make_id = $2, model_name = $3
		WHERE model_id = $1
	`

	result, err := r.db.Pool.Exec(ctx, query, model.ModelID, model.MakeID, model.ModelName)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("model not found")
	}

	return nil
}

func (r *PostgresVehicleModelRepository) DeleteModel(ctx context.Context, id int) error {
	query := `DELETE FROM vehicle_models WHERE model_id = $1`

	result, err := r.db.Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("model not found")
	}

	return nil
}

func (r *PostgresVehicleModelRepository) GetSubmodelsByModel(ctx context.Context, modelID int) ([]*vehicleSubmodels.VehicleSubmodel, error) {
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

	var submodels []*vehicleSubmodels.VehicleSubmodel
	for rows.Next() {
		submodel := &vehicleSubmodels.VehicleSubmodel{}
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
