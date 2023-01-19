package repo

import (
	"context"

	"github.com/joyzem/proxy-project/services/product/domain"
)

// Репозиторий обращается к источнику данных. В данном случае к БД
type ProductRepo interface {
	CreateProduct(context.Context, domain.Product) (*domain.Product, error)
	GetProducts(context.Context) ([]domain.Product, error)
	UpdateProduct(context.Context, domain.Product) (*domain.Product, error)
	DeleteProduct(context.Context, int64) error
}
