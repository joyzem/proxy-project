package product

import (
	"context"
)

// Весь сервис товаров для внешнего источника представлен этим интерфейсом
type Service interface {
	// Создать товар
	CreateProduct(ctx context.Context, name string, price int32, unit Unit) (*Product, error)
	// Получить все товары
	GetProducts(context.Context) ([]Product, error)
	// Обновить информацию о товаре
	UpdateProduct(context.Context, Product) (*Product, error)
	// Удалить товар
	DeleteProduct(context.Context, int64) error
	// Создать единицы измерения
	CreateUnit(context.Context, string) (*Unit, error)
	// Получить все единицы измерения
	GetUnits(context.Context) ([]Unit, error)
	// Обновить единицу измерения
	UpdateUnit(context.Context, Unit) (*Unit, error)
	// Удалить единицу измерения
	DeleteUnit(context.Context, int64) error
}

func NewService() {

}
