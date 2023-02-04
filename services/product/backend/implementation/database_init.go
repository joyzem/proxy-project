package implementation

import (
	"database/sql"
)

// Функция инициализации таблиц
func InitDatabase(db sql.DB) error {
	// запрос на создание таблицы units
	createUnitsTableSql := `
		CREATE TABLE IF NOT EXISTS units (
			id serial PRIMARY KEY,
			name text NOT NULL
		);
	`
	_, err := db.Exec(createUnitsTableSql)
	if err != nil {
		return err
	}
	// добавление значения по умолчанию
	insertDefaultValue := `
		INSERT INTO units (id, name)
		VALUES (0, 'Не указано')
		ON CONFLICT DO NOTHING
	`
	_, err = db.Exec(insertDefaultValue)
	if err != nil {
		return err
	}
	// запрос на создание таблицы products
	createProductsTableSql := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			price INTEGER NOT NULL,
			unit_id INTEGER NOT NULL DEFAULT 0,
			FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE SET DEFAULT ON UPDATE CASCADE
		)
	`
	_, err = db.Exec(createProductsTableSql)
	if err != nil {
		return err
	}
	return nil
}
