package product

import "context"

// Единицы измерения
type Unit struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Репозиторий
type UnitRepo interface {
	CreateUnit(context.Context, string) (*Unit, error)
	GetUnits(context.Context) ([]Unit, error)
	UpdateUnit(context.Context, Unit) (*Unit, error)
	DeleteUnit(context.Context, int64) error
}
