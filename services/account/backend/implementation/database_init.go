package implementation

import "database/sql"

func InitDatabase(db sql.DB) error {
	createAccountsTable := `
		CREATE TABLE IF NOT EXISTS accounts (
			id SERIAL PRIMARY KEY,
			bank_name TEXT NOT NULL,
			bank_identity_number CHAR(9) NOT NULL
		);
	`
	_, err := db.Exec(createAccountsTable)
	if err != nil {
		return err
	}
	return nil
}
