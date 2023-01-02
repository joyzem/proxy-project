package product

import (
	"context"

	"github.com/joyzem/proxy-project/internal/services/unit"
)

// Описание структуры "Товар"
type Product struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Price int32     `json:"price"`
	Unit  unit.Unit `json:"unit"` // Единицы измерения
}

// Репозиторий обращается к источнику данных. В данном случае к БД
type Repository interface {
	CreateProduct(context.Context, Product) error
	GetProducts(context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id int) (Product, error)
	UpdateProduct(ctx context.Context, id int, product Product) error
	DeleteProduct(ctx context.Context, id int) error
}
