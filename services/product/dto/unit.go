package dto

import "github.com/joyzem/proxy-project/services/product/domain"

type (
	CreateUnitRequest struct {
		Unit string `json:"unit"`
	}
	CreateUnitResponse struct {
		Unit *domain.Unit `json:"unit,omitempty"`
		Err  string       `json:"error,omitempty"`
	}
	GetUnitsRequest struct {
	}
	GetUnitsResponse struct {
		Units []domain.Unit `json:"units,omitempty"`
		Err   string        `json:"error,omitempty"`
	}
	UpdateUnitRequest struct {
		Unit domain.Unit `json:"unit"`
	}
	UpdateUnitResponse struct {
		Unit *domain.Unit `json:"unit,omitempty"`
		Err  string       `json:"error,omitempty"`
	}
	DeleteUnitRequest struct {
		Id int `json:"id"`
	}
	DeleteUnitResponse struct {
		Err string `json:"error,omitempty"`
	}
)
