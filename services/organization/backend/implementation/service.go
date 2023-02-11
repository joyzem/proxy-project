package implementation

import (
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/organization/backend/repo"
	svc "github.com/joyzem/proxy-project/services/organization/backend/service"
	"github.com/joyzem/proxy-project/services/organization/domain"
)

type service struct {
	organizationRepo repo.OrganizationRepo
}

func NewService(organizationRepo repo.OrganizationRepo) svc.OrganizationService {
	return &service{
		organizationRepo: organizationRepo,
	}
}

func (s *service) GetOrganizations() ([]domain.Organization, error) {
	organizations, err := s.organizationRepo.GetOrganizations()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return organizations, nil
}

func (s *service) CreateOrganization(name string, address string, accountId int, chief string, financialChief string) (*domain.Organization, error) {
	org := domain.Organization{Name: name, Address: address, AccountId: accountId, Chief: chief, FinancialChief: financialChief}
	createdOrganization, err := s.organizationRepo.CreateOrganization(org)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdOrganization, nil
}

func (s *service) UpdateOrganization(newOrganization domain.Organization) (*domain.Organization, error) {
	updatedOrganization, err := s.organizationRepo.UpdateOrganization(newOrganization)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedOrganization, nil
}

func (s *service) DeleteOrganization(id int) error {
	if err := s.organizationRepo.DeleteOrganization(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

func (s *service) OrganizationById(id int) (*domain.Organization, error) {
	org, err := s.organizationRepo.OrganizationById(id)
	if err != nil {
		base.LogError(err)
	}
	return org, err
}
