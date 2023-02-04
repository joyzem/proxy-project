package implementation

import (
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/backend/repo"
	svc "github.com/joyzem/proxy-project/services/product/backend/service"
	"github.com/joyzem/proxy-project/services/product/domain"
)

type service struct {
	productRepo repo.ProductRepo
	unitRepo    repo.UnitRepo
}

func NewService(productRepo repo.ProductRepo, unitRepo repo.UnitRepo) svc.Service {
	return &service{
		productRepo: productRepo,
		unitRepo:    unitRepo,
	}
}

func (s *service) CreateProduct(name string, price int, unitId int) (*domain.Product, error) {
	p := domain.Product{Name: name, Price: price, Unit: domain.Unit{Id: unitId}}
	createdProduct, err := s.productRepo.CreateProduct(p)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdProduct, nil
}

func (s *service) GetProducts() ([]domain.Product, error) {
	products, err := s.productRepo.GetProducts()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return products, nil
}

func (s *service) UpdateProduct(oldProduct domain.Product) (*domain.Product, error) {
	updatedProduct, err := s.productRepo.UpdateProduct(oldProduct)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedProduct, nil
}

func (s *service) DeleteProduct(id int) error {
	if err := s.productRepo.DeleteProduct(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

func (s *service) CreateUnit(unit string) (*domain.Unit, error) {
	createdUnit, err := s.unitRepo.CreateUnit(unit)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdUnit, nil
}

func (s *service) GetUnits() ([]domain.Unit, error) {
	units, err := s.unitRepo.GetUnits()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return units, nil
}

func (s *service) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	updatedUnit, err := s.unitRepo.UpdateUnit(unit)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedUnit, nil
}

func (s *service) DeleteUnit(id int) error {
	if err := s.unitRepo.DeleteUnit(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}
