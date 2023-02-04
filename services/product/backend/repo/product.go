package repo

import (
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Репозиторий обращается к источнику данных. В данном случае к БД
type ProductRepo interface {
	CreateProduct(domain.Product) (*domain.Product, error)
	GetProducts() ([]domain.Product, error)
	UpdateProduct(domain.Product) (*domain.Product, error)
	DeleteProduct(int) error
}
