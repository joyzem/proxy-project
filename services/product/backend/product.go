package product

import (
	"context"
)

// Структура "Товар"
type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
	Unit  Unit   `json:"unit"` // Единицы измерения
}

// Репозиторий обращается к источнику данных. В данном случае к БД
type ProductRepo interface {
	CreateProduct(context.Context, Product) (*Product, error)
	GetProducts(context.Context) ([]Product, error)
	UpdateProduct(context.Context, Product) (*Product, error)
	DeleteProduct(context.Context, int64) error
}
