package implementation

import (
	"database/sql"

	"github.com/joyzem/proxy-project/services/account/backend/repo"
	"github.com/joyzem/proxy-project/services/account/domain"
)

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) repo.AccountRepo {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateAccount(acc domain.Account) (*domain.Account, error) {
	sql := `
		INSERT INTO accounts (bank_name, bank_identity_number)
		VALUES ($1, $2) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(acc.BankName, acc.BankIdentityNumber).Scan(&insertedId); err != nil {
		return nil, err
	}

	return r.AccountById(insertedId)
}

func (r *repository) GetAccounts() ([]domain.Account, error) {
	sql := `
		SELECT * FROM accounts ORDER BY bank_name ASC
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return []domain.Account{}, err
	}
	defer rows.Close()

	accounts := []domain.Account{}

	for rows.Next() {
		account := domain.Account{} // Current Product
		rows.Scan(&account.Id, &account.BankName, &account.BankIdentityNumber)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *repository) UpdateAccount(acc domain.Account) (*domain.Account, error) {
	sql := `
		UPDATE accounts 
			SET bank_name = $1, bank_identity_number = $2 
		WHERE id = $3
	`
	rows, err := r.db.Query(sql, acc.BankName, acc.BankIdentityNumber, acc.Id)
	if err != nil {
		return nil, err
	}

	rows.Close()

	return r.AccountById(acc.Id)
}

func (r *repository) DeleteAccount(id int) error {
	sql := `
		DELETE FROM accounts WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *repository) AccountById(id int) (*domain.Account, error) {
	sql := `
		SELECT * FROM accounts WHERE id = $1
	`
	account := domain.Account{}
	if err := r.db.QueryRow(sql, id).Scan(
		&account.Id,
		&account.BankName,
		&account.BankIdentityNumber,
	); err != nil {
		return nil, err
	}
	return &account, nil
}
