package repo

import (
	"context"

	"github.com/joyzem/proxy-project/services/product/domain"
)

// Репозиторий
type UnitRepo interface {
	CreateUnit(context.Context, string) (*domain.Unit, error)
	GetUnits(context.Context) ([]domain.Unit, error)
	UpdateUnit(context.Context, domain.Unit) (*domain.Unit, error)
	DeleteUnit(context.Context, int64) error
}
