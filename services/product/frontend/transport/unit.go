package transport

import "github.com/joyzem/proxy-project/services/product/domain"

type GetUnitsResponse struct {
	Units []domain.Unit `json:"units"`
	Err   error         `json:"error"`
}

type DeleteUnitRequest struct {
	Id int `json:"id"`
}

type DeleteUnitResponse struct {
	Err error `json:"error"`
}

type CreateUnitRequest struct {
	Unit string `json:"unit"`
}

type CreateUnitResponse struct {
	Unit *domain.Unit `json:"unit"`
	Err  error        `json:"error"`
}

type UpdateUnitRequest struct {
	Unit *domain.Unit `json:"unit"`
}
