package implementation

import (
	"context"
	"database/sql"
	"sort"

	"github.com/go-kit/log"
	repo "github.com/joyzem/proxy-project/services/product/backend/repo"
	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/utils"
)

type unitRepository struct {
	db     *sql.DB
	logger log.Logger
}

func NewUnitRepository(db *sql.DB) repo.UnitRepo {
	return &unitRepository{
		db: db,
	}
}

func (r *unitRepository) CreateUnit(ctx context.Context, unit string) (*domain.Unit, error) {
	sql := `
			INSERT INTO units (name) VALUES ($1) RETURNING id
	`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	defer result.Close()

	var insertedId int64
	if err := result.QueryRow(unit).Scan(&insertedId); err != nil {
		utils.LogError(err)
		return nil, err
	}
	insertedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	createdUnit := domain.Unit{}
	if err := r.db.QueryRowContext(ctx, insertedUnitSql, insertedId).Scan(&createdUnit.Id, &createdUnit.Name); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return &createdUnit, nil
}

func (r *unitRepository) GetUnits(ctx context.Context) ([]domain.Unit, error) {
	sql := `
			SELECT * FROM units		
	`
	rows, err := r.db.QueryContext(ctx, sql)
	if err != nil {
		utils.LogError(err)
		return []domain.Unit{}, err
	}
	defer rows.Close()

	units := []domain.Unit{}

	for rows.Next() {
		unit := domain.Unit{}
		err := rows.Scan(&unit.Id, &unit.Name)
		if err != nil {
			utils.LogError(err)
			continue
		}
		units = append(units, unit)
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].Id < units[j].Id
	})
	return units, nil
}

func (r *unitRepository) UpdateUnit(ctx context.Context, unit domain.Unit) (*domain.Unit, error) {
	sql := `
			UPDATE units SET name = $1 WHERE id = $2 RETURNING id
	`
	result, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	defer result.Close()

	var updatedId int
	if err := result.QueryRowContext(ctx, unit.Name, unit.Id).Scan(&updatedId); err != nil {
		utils.LogError(err)
		return nil, err
	}
	updatedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	updatedUnit := domain.Unit{}
	if err := r.db.QueryRowContext(ctx, updatedUnitSql, updatedId).Scan(&updatedUnit.Id, &updatedUnit.Name); err != nil {
		utils.LogError(err)
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
		utils.LogError(err)
		return err
	}
	return nil
}
