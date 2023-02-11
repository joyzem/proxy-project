package implementation

import "database/sql"

func InitDatabase(db sql.DB) error {
	createCustomersTableSql := `
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		)
	`
	_, err := db.Exec(createCustomersTableSql)
	return err
}
