package implementation

import (
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/backend/repo"
	svc "github.com/joyzem/proxy-project/services/product/backend/service"
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Реализация сервиса
type service struct {
	productRepo repo.ProductRepo
	unitRepo    repo.UnitRepo
}

// Возвращает реализацию сервиса
func NewService(productRepo repo.ProductRepo, unitRepo repo.UnitRepo) svc.Service {
	return &service{
		productRepo: productRepo,
		unitRepo:    unitRepo,
	}
}

// CreateProduct - создает новый товар с переданными параметрами (name, price, unitId)
func (s *service) CreateProduct(name string, price int, unitId int) (*domain.Product, error) {
	p := domain.Product{Name: name, Price: price, UnitId: unitId}
	createdProduct, err := s.productRepo.CreateProduct(p)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdProduct, nil
}

// GetProducts - возвращает все товары
func (s *service) GetProducts() ([]domain.Product, error) {
	products, err := s.productRepo.GetProducts()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return products, nil
}

// UpdateProduct - обновляет информацию о товаре
func (s *service) UpdateProduct(newProduct domain.Product) (*domain.Product, error) {
	updatedProduct, err := s.productRepo.UpdateProduct(newProduct)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedProduct, nil
}

// DeleteProduct - удаляет товар по id
func (s *service) DeleteProduct(id int) error {
	if err := s.productRepo.DeleteProduct(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

// ProductById - возвращает товар по идентификатору
func (s *service) ProductById(id int) (*domain.Product, error) {
	product, err := s.productRepo.ProductById(id)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return product, nil
}

// CreateUnit - добавляет единицу измерения.
func (s *service) CreateUnit(unit string) (*domain.Unit, error) {
	createdUnit, err := s.unitRepo.CreateUnit(unit)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return createdUnit, nil
}

// GetUnits - возвращает список всех единиц измерения.
func (s *service) GetUnits() ([]domain.Unit, error) {
	units, err := s.unitRepo.GetUnits()
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return units, nil
}

// UpdateUnit - обновляет информацию об единице измерения.
func (s *service) UpdateUnit(unit domain.Unit) (*domain.Unit, error) {
	updatedUnit, err := s.unitRepo.UpdateUnit(unit)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return updatedUnit, nil
}

// DeleteUnit - удаляет единицу измерения.
func (s *service) DeleteUnit(id int) error {
	if err := s.unitRepo.DeleteUnit(id); err != nil {
		base.LogError(err)
		return err
	}
	return nil
}

// UnitById - возвращает единицу измерения по идентификатору
func (s *service) UnitById(id int) (*domain.Unit, error) {
	unit, err := s.unitRepo.UnitById(id)
	if err != nil {
		base.LogError(err)
		return nil, err
	}
	return unit, nil
}
