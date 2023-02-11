package implementation

import (
	"github.com/joyzem/proxy-project/services/account/backend/repo"
	svc "github.com/joyzem/proxy-project/services/account/backend/service"
	"github.com/joyzem/proxy-project/services/account/domain"
	"github.com/joyzem/proxy-project/services/base"
)

type service struct {
	accountRepo repo.AccountRepo
}

func NewService(accountRepo repo.AccountRepo) svc.Service {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) CreateAccount(bankName string, bankIdentityCode string) (*domain.Account, error) {
	acc := domain.Account{BankName: bankName, BankIdentityNumber: bankIdentityCode}
	createdAccount, err := s.accountRepo.CreateAccount(acc)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdAccount, nil
}

func (s *service) GetAccounts() ([]domain.Account, error) {
	accounts, err := s.accountRepo.GetAccounts()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return accounts, nil
}

func (s *service) UpdateAccount(account domain.Account) (*domain.Account, error) {
	updatedAccount, err := s.accountRepo.UpdateAccount(account)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedAccount, nil
}

func (s *service) DeleteAccount(id int) error {
	if err := s.accountRepo.DeleteAccount(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

func (s *service) AccountById(id int) (*domain.Account, error) {
	acc, err := s.accountRepo.AccountById(id)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return acc, nil
}
