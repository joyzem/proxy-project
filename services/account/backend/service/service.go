package service

import "github.com/joyzem/proxy-project/services/account/domain"

type Service interface {
	CreateAccount(bankName string, bankIdentityCode string) (*domain.Account, error)
	GetAccounts() ([]domain.Account, error)
	UpdateAccount(domain.Account) (*domain.Account, error)
	DeleteAccount(int) error
	AccountById(int) (*domain.Account, error)
}
