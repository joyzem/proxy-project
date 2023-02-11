package implementation

import "database/sql"

func InitDatabase(db *sql.DB) error {
	createEmployeesTableSql := `
		CREATE TABLE IF NOT EXISTS employees (
			id SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			middle_name TEXT NOT NULL,
			post TEXT NOT NULL,
			passport_series CHAR(4) NOT NULL,
			passport_number CHAR(6) NOT NULL,
			passport_issued_by TEXT NOT NULL,
			passport_date_of_issue DATE NOT NULL
		)
	`
	_, err := db.Exec(createEmployeesTableSql)
	return err
}
