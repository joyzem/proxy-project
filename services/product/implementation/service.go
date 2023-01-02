package implementation

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joyzem/proxy-project/internal/services/product"
)

type service struct {
	repository product.Repository
	log        log.Logger
}

func NewService(rep product.Repository, logger log.Logger) product.Service {
	return &service{
		repository: rep,
		log:        logger,
	}
}

func (s *service) CreateProduct(ctx context.Context, product product.Product) error {
	logger := log.With(s.log, "Method", "Create")
	if err := s.repository.CreateProduct(ctx, product); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

func (s *service) GetProducts(ctx context.Context) ([]product.Product, error) {
	logger := log.With(s.log, "Method", "GetProducts")
	products, err := s.repository.GetProducts(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	return products, nil
}
func (s *service) GetProduct(ctx context.Context, id int) (product.Product, error) {
	logger := log.With(s.log, "Method", "GetProduct")
	product, err := s.repository.GetProduct(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return product, err
	}
	return product, nil
}
func (s *service) UpdateProduct(ctx context.Context, id int, product product.Product) error {
	logger := log.With(s.log, "Method", "Update")
	if err := s.repository.UpdateProduct(ctx, id, product); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}
func (s *service) DeleteProduct(ctx context.Context, id int) error {
	logger := log.With(s.log, "Method", "Delete")
	if err := s.repository.DeleteProduct(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}
