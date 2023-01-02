package implementation

import (
	"context"

	"github.com/joyzem/proxy-project/internal/services/product"
)

type service struct {
	repository product.Repository
}

func NewService(rep product.Repository) product.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) CreateProduct(ctx context.Context, product product.Product) error {
	if err := s.repository.CreateProduct(ctx, product); err != nil {
		return err
	}
	return nil
}

func (s *service) GetProducts(ctx context.Context) ([]product.Product, error) {
	products, err := s.repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *service) GetProduct(ctx context.Context, id int) (product.Product, error) {
	product, err := s.repository.GetProduct(ctx, id)
	if err != nil {
		return product, err
	}
	return product, nil
}
func (s *service) UpdateProduct(ctx context.Context, id int, product product.Product) error {
	return s.repository.UpdateProduct(ctx, id, product)
}
func (s *service) DeleteProduct(ctx context.Context, id int) error {
	return s.repository.DeleteProduct(ctx, id)
}
