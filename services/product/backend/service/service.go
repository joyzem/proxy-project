package service

import (
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Весь сервис товаров для внешнего источника представлен этим интерфейсом
type Service interface {
	// Создать товар
	CreateProduct(name string, price int, unitId int) (*domain.Product, error)
	// Получить все товары
	GetProducts() ([]domain.Product, error)
	// Обновить информацию о товаре
	UpdateProduct(domain.Product) (*domain.Product, error)
	// Удалить товар
	DeleteProduct(int) error
	// Создать единицы измерения
	CreateUnit(string) (*domain.Unit, error)
	// Получить все единицы измерения
	GetUnits() ([]domain.Unit, error)
	// Обновить единицу измерения
	UpdateUnit(domain.Unit) (*domain.Unit, error)
	// Удалить единицу измерения
	DeleteUnit(int) error
}
