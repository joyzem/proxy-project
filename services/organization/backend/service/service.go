package service

import "github.com/joyzem/proxy-project/services/organization/domain"

type Service interface {
	GetOrganizations() ([]domain.Organization, error)
	CreateOrganization(domain.Organization) (*domain.Organization, error)
	UpdateOrganization(domain.Organization) (*domain.Organization, error)
	DeleteOrganization(int) error
}
