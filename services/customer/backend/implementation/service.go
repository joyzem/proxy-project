package implementation

import (
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/customer/backend/repo"
	svc "github.com/joyzem/proxy-project/services/customer/backend/service"
	"github.com/joyzem/proxy-project/services/customer/domain"
)

type service struct {
	repo repo.CustomerRepo
}

func NewService(repo repo.CustomerRepo) svc.CustomerService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateCustomer(name string) (*domain.Customer, error) {
	customer := domain.Customer{
		Name: name,
	}
	result, err := s.repo.CreateCustomer(customer)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return result, nil
}

func (s *service) GetCustomers() ([]domain.Customer, error) {
	result, err := s.repo.GetCustomers()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return result, nil
}

func (s *service) UpdateCustomer(customer domain.Customer) (*domain.Customer, error) {
	result, err := s.repo.UpdateCustomer(customer)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return result, nil
}

func (s *service) DeleteCustomer(id int) error {
	err := s.repo.DeleteCustomer(id)
	if err != nil {
		base.LogError(err)
	}
	return err
}

func (s *service) CustomerById(id int) (*domain.Customer, error) {
	customer, err := s.repo.CustomerById(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
