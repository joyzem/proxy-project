package implementation

import (
	"database/sql"
	"sort"

	repo "github.com/joyzem/proxy-project/services/product/backend/repo"
	"github.com/joyzem/proxy-project/services/product/domain"
)

type unitRepository struct {
	db *sql.DB
}

func NewUnitRepository(db *sql.DB) repo.UnitRepo {
	return &unitRepository{
		db: db,
	}
}

func (r *unitRepository) CreateUnit(unit string) (*domain.Unit, error) {
	sql := `
			INSERT INTO units (name) VALUES ($1) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int64
	if err := result.QueryRow(unit).Scan(&insertedId); err != nil {
		return nil, err
	}
	insertedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	createdUnit := domain.Unit{}
	if err := r.db.QueryRow(insertedUnitSql, insertedId).Scan(&createdUnit.Id, &createdUnit.Name); err != nil {
		return nil, err
	}
	return &createdUnit, nil
}

func (r *unitRepository) GetUnits() ([]domain.Unit, error) {
	sql := `
			SELECT * FROM units		
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	units := []domain.Unit{}

	for rows.Next() {
		unit := domain.Unit{}
		rows.Scan(&unit.Id, &unit.Name)
		units = append(units, unit)
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].Id < units[j].Id
	})
	return units, nil
}

func (r *unitRepository) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	sql := `
			UPDATE units SET name = $1 WHERE id = $2 RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var updatedId int
	if err := result.QueryRow(unit.Name, unit.Id).Scan(&updatedId); err != nil {
		return nil, err
	}
	updatedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`
	updatedUnit := domain.Unit{}
	if err := r.db.QueryRow(updatedUnitSql, updatedId).Scan(&updatedUnit.Id, &updatedUnit.Name); err != nil {
		return nil, err
	}
	return &updatedUnit, nil
}

func (r *unitRepository) DeleteUnit(id int) error {
	sql := `
			DELETE FROM units WHERE id = $1		
	`
	_, err := r.db.Exec(sql, id)
	return err
}
