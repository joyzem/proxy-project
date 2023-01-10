package implementation

import (
	"context"
	"database/sql"

	"github.com/go-kit/log"
	product "github.com/joyzem/proxy-project/services/product/backend"
	"github.com/joyzem/proxy-project/services/utils"
)

type unitRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewUnitRepository(db *sql.DB, logger log.Logger) (product.UnitRepo, error) {
	return &unitRepository{
		db:     db,
		logger: log.With(logger, "unit repo"),
	}, nil
}

func (r *unitRepository) CreateUnit(ctx context.Context, unit string) (*product.Unit, error) {
	sql := `
			INSERT INTO units (name) VALUES ($1) RETURNING id
	`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	defer result.Close()

	var insertedId int64
	if err := result.QueryRow(unit).Scan(&insertedId); err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	insertedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	createdUnit := product.Unit{}
	if err := r.db.QueryRowContext(ctx, insertedUnitSql, insertedId).Scan(&createdUnit.Id, &createdUnit.Name); err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	return &createdUnit, nil
}

func (r *unitRepository) GetUnits(ctx context.Context) ([]product.Unit, error) {
	sql := `
			SELECT * FROM units		
	`
	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return []product.Unit{}, err
	}
	defer rows.Close()

	units := []product.Unit{}

	for rows.Next() {
		unit := product.Unit{}
		err := rows.Scan(&unit.Id, &unit.Name)
		if err != nil {
			utils.LogError(&r.logger, err)
			continue
		}
		units = append(units, unit)
	}

	return units, nil
}

func (r *unitRepository) UpdateUnit(ctx context.Context, unit product.Unit) (*product.Unit, error) {
	sql := `
			UPDATE units SET name = $1 WHERE id = $2 RETURNING id
	`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	defer result.Close()

	var updatedId int
	if err := result.QueryRowContext(ctx, unit.Name, unit.Id).Scan(&updatedId); err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	updatedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	updatedUnit := product.Unit{}
	if err := r.db.QueryRowContext(ctx, updatedUnitSql, updatedId).Scan(&updatedUnit.Id, &updatedUnit.Name); err != nil {
		utils.LogError(&r.logger, err)
		return nil, err
	}
	return &updatedUnit, nil
}

func (r *unitRepository) DeleteUnit(ctx context.Context, id int64) error {
	sql := `
			DELETE FROM units WHERE id = $1		
	`
	_, err := r.db.ExecContext(ctx, sql, id)
	if err != nil {
		utils.LogError(&r.logger, err)
		return err
	}
	return nil
}
