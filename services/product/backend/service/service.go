package service

import (
	"context"

	"github.com/joyzem/proxy-project/services/product/domain"
)

// Весь сервис товаров для внешнего источника представлен этим интерфейсом
type Service interface {
	// Создать товар
	CreateProduct(ctx context.Context, name string, price int, unitId int) (*domain.Product, error)
	// Получить все товары
	GetProducts(context.Context) ([]domain.Product, error)
	// Обновить информацию о товаре
	UpdateProduct(context.Context, domain.Product) (*domain.Product, error)
	// Удалить товар
	DeleteProduct(context.Context, int64) error
	// Создать единицы измерения
	CreateUnit(context.Context, string) (*domain.Unit, error)
	// Получить все единицы измерения
	GetUnits(context.Context) ([]domain.Unit, error)
	// Обновить единицу измерения
	UpdateUnit(context.Context, domain.Unit) (*domain.Unit, error)
	// Удалить единицу измерения
	DeleteUnit(context.Context, int64) error
}
