package product

import "context"

// Весь сервис товаров для внешнего мира представлен этим интерфейсом
type Service interface {
	CreateProduct(ctx context.Context, product Product) error
	GetProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id int) (Product, error)
	UpdateProduct(ctx context.Context, id int, product Product) error
	DeleteProduct(ctx context.Context, id int) error
}
