package implementation

import (
	"context"

	"github.com/joyzem/proxy-project/services/product/backend/repo"
	svc "github.com/joyzem/proxy-project/services/product/backend/service"
	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/utils"
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

func (s *service) CreateProduct(ctx context.Context, name string, price int, unitId int) (*domain.Product, error) {
	p := domain.Product{Name: name, Price: price, Unit: domain.Unit{Id: unitId}}
	createdProduct, err := s.productRepo.CreateProduct(ctx, p)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return createdProduct, nil
}

func (s *service) GetProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := s.productRepo.GetProducts(ctx)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, oldProduct domain.Product) (*domain.Product, error) {
	updatedProduct, err := s.productRepo.UpdateProduct(ctx, oldProduct)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return updatedProduct, nil
}

func (s *service) DeleteProduct(ctx context.Context, id int64) error {
	if err := s.productRepo.DeleteProduct(ctx, id); err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

func (s *service) CreateUnit(ctx context.Context, unit string) (*domain.Unit, error) {
	createdUnit, err := s.unitRepo.CreateUnit(ctx, unit)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return createdUnit, nil
}

func (s *service) GetUnits(ctx context.Context) ([]domain.Unit, error) {
	units, err := s.unitRepo.GetUnits(ctx)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return units, nil
}

func (s *service) UpdateUnit(ctx context.Context, unit domain.Unit) (*domain.Unit, error) {
	updatedUnit, err := s.unitRepo.UpdateUnit(ctx, unit)
	if err != nil {
		utils.LogError(err)
		return nil, err
	}
	return updatedUnit, nil
}

func (s *service) DeleteUnit(ctx context.Context, id int64) error {
	if err := s.unitRepo.DeleteUnit(ctx, id); err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}
