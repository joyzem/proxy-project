package repo

import (
	"github.com/joyzem/proxy-project/services/product/domain"
)

// Репозиторий
type UnitRepo interface {
	CreateUnit(string) (*domain.Unit, error)
	GetUnits() ([]domain.Unit, error)
	UpdateUnit(domain.Unit) (*domain.Unit, error)
	DeleteUnit(int) error
}
