package implementation

import "database/sql"

func InitDatabase(db sql.DB) error {
	createOrganizationsTableSql := `
		CREATE TABLE IF NOT EXISTS organizations (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			address TEXT NOT NULL,
			account_id INTEGER NOT NULL,
			chief TEXT NOT NULL,
			financial_chief TEXT NOT NULL,
			FOREIGN KEY (account_id) REFERENCES accounts(id) ON DELETE RESTRICT ON UPDATE CASCADE
		)
	`
	_, err := db.Exec(createOrganizationsTableSql)
	return err
}
