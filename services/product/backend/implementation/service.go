package implementation

import (
	"context"

	"github.com/go-kit/log"
	product "github.com/joyzem/proxy-project/services/product/backend"
	"github.com/joyzem/proxy-project/services/utils"
)

type service struct {
	productRepo product.ProductRepo
	unitRepo    product.UnitRepo
	log         log.Logger
}

func NewService(productRepo product.ProductRepo, unitRepo product.UnitRepo, logger log.Logger) product.Service {
	return &service{
		productRepo: productRepo,
		unitRepo:    unitRepo,
		log:         logger,
	}
}

func (s *service) CreateProduct(ctx context.Context, name string, price int32, unit product.Unit) (*product.Product, error) {
	logger := log.With(s.log, "Method", "Create Product")
	p := product.Product{Name: name, Price: price, Unit: unit}
	createdProduct, err := s.productRepo.CreateProduct(ctx, p)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return createdProduct, nil
}

func (s *service) GetProducts(ctx context.Context) ([]product.Product, error) {
	logger := log.With(s.log, "Method", "Get Products")
	products, err := s.productRepo.GetProducts(ctx)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return products, nil
}

func (s *service) UpdateProduct(ctx context.Context, oldProduct product.Product) (*product.Product, error) {
	logger := log.With(s.log, "Method", "Update Product")
	updatedProduct, err := s.productRepo.UpdateProduct(ctx, oldProduct)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return updatedProduct, nil
}

func (s *service) DeleteProduct(ctx context.Context, id int64) error {
	logger := log.With(s.log, "Method", "Delete Product")
	if err := s.productRepo.DeleteProduct(ctx, id); err != nil {
		utils.LogError(&logger, err)
		return err
	}
	return nil
}

func (s *service) CreateUnit(ctx context.Context, unit string) (*product.Unit, error) {
	logger := log.With(s.log, "Method", "Create Unit")
	createdUnit, err := s.unitRepo.CreateUnit(ctx, unit)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return createdUnit, nil
}

func (s *service) GetUnits(ctx context.Context) ([]product.Unit, error) {
	logger := log.With(s.log, "Method", "Get Units")
	units, err := s.unitRepo.GetUnits(ctx)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return units, nil
}

func (s *service) UpdateUnit(ctx context.Context, unit product.Unit) (*product.Unit, error) {
	logger := log.With(s.log, "Method", "Update Unit")
	updatedUnit, err := s.unitRepo.UpdateUnit(ctx, unit)
	if err != nil {
		utils.LogError(&logger, err)
		return nil, err
	}
	return updatedUnit, nil
}

func (s *service) DeleteUnit(ctx context.Context, id int64) error {
	logger := log.With(s.log, "Method", "Delete Unit")
	if err := s.unitRepo.DeleteUnit(ctx, id); err != nil {
		utils.LogError(&logger, err)
		return err
	}
	return nil
}
