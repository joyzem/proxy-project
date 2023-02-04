package implementation

import (
	"database/sql"
	"sort"

	repo "github.com/joyzem/proxy-project/services/product/backend/repo"
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Реализация репозитория единиц измерения
type unitRepository struct {
	db *sql.DB
}

// Возвращает репозиторий единиц измерения
func NewUnitRepository(db *sql.DB) repo.UnitRepo {
	return &unitRepository{
		db: db,
	}
}

// CreateUnit добавляет новую единицу измерения в базу данных и возвращает указатель на созданную единицу измерения или ошибку, если она возникает.
func (r *unitRepository) CreateUnit(unit string) (*domain.Unit, error) {
	// Запрос на добавление единицы измерения с возвратом id
	sql := `
			INSERT INTO units (name) VALUES ($1) RETURNING id
	`
	// Подготовка запроса
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Выполнение запроса и сканирование результата в insertedId
	var insertedId int64
	if err := result.QueryRow(unit).Scan(&insertedId); err != nil {
		return nil, err
	}

	// SQL-запрос для получения данных новой единицы измерения
	insertedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`

	// Инициализация переменной с новой единицей измерения
	createdUnit := domain.Unit{}

	// Выполнение запроса и сканирование результата
	if err := r.db.QueryRow(insertedUnitSql, insertedId).Scan(&createdUnit.Id, &createdUnit.Name); err != nil {
		return nil, err
	}

	return &createdUnit, nil
}

// GetUnits возвращает список всех единиц измерения из базы данных или ошибку, если она возникает.
func (r *unitRepository) GetUnits() ([]domain.Unit, error) {
	// Формирование SQL запроса для получения всех единиц измерения из таблицы units
	sql := `SELECT * FROM units`
	// Выполнение запроса
	rows, err := r.db.Query(sql)
	if err != nil {
		// Возврат ошибки, если произошла ошибка при выполнении запроса
		return nil, err
	}
	// Освобождение ресурсов, занятых результатом запроса
	defer rows.Close()

	// Инициализация массива единиц измерения
	units := []domain.Unit{}

	// Обработка каждой строки результата запроса
	for rows.Next() {
		// Инициализация структуры единицы измерения
		unit := domain.Unit{}
		// Сканирование данных из текущей строки результата запроса
		rows.Scan(&unit.Id, &unit.Name)
		// Добавление единицы измерения в массив
		units = append(units, unit)
	}

	// Сортировка по id
	sort.Slice(units, func(i, j int) bool {
		return units[i].Id < units[j].Id
	})

	return units, nil
}

// UpdateUnit обновляет единицу измерения в базе данных.
// unit - структура с новыми значениями единицы измерения.
// Возвращает указатель на обновленную единицу измерения или ошибку, если она возникла.
func (r *unitRepository) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	// SQL-запрос для обновления информации об единице измерения
	sql := `
			UPDATE units SET name = $1 WHERE id = $2 RETURNING id
	`
	// Подготовка запроса
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	// Выполнение запроса и сканирование результата (ID)
	var updatedId int
	if err := result.QueryRow(unit.Name, unit.Id).Scan(&updatedId); err != nil {
		return nil, err
	}

	// SQL-запрос для получения данных обновленной единицы измерения
	updatedUnitSql := `
			SELECT * FROM units WHERE id = $1	
	`

	// Инициализация переменной с обновленной единицей измерения
	updatedUnit := domain.Unit{}
	if err := r.db.QueryRow(updatedUnitSql, updatedId).Scan(&updatedUnit.Id, &updatedUnit.Name); err != nil {
		return nil, err
	}

	return &updatedUnit, nil
}

// DeleteUnit выполняет удаление единицы измерения из базы данных по её ID.
// ID единицы измерения передается в качестве параметра.
func (r *unitRepository) DeleteUnit(id int) error {
	sql := `
			DELETE FROM units WHERE id = $1		
	`
	_, err := r.db.Exec(sql, id)
	return err
}
