package transport

import (
	"github.com/joyzem/proxy-project/services/product/domain"
)

type (
	CreateProductRequest struct {
		Name   string `json:"name"`
		Price  int    `json:"price"`
		UnitId int    `json:"unit_id"`
	}
	CreateProductResponse struct {
		Product *domain.Product `json:"product"`
		Err     error           `json:"error"`
	}
	GetProductsRequest struct {
	}
	GetProductsResponse struct {
		Products []domain.Product `json:"products"`
		Err      error            `json:"error"`
	}
	UpdateProductRequest struct {
		Product domain.Product `json:"product"`
	}
	UpdateProductResponse struct {
		Product *domain.Product `json:"product"`
		Err     error           `json:"error"`
	}
	DeleteProductRequest struct {
		Id int64 `json:"id"`
	}
	DeleteProductResponse struct {
		Err error `json:"error"`
	}
	CreateUnitRequest struct {
		Unit string `json:"unit"`
	}
	CreateUnitResponse struct {
		Unit *domain.Unit `json:"unit"`
		Err  error        `json:"error"`
	}
	GetUnitsRequest struct {
	}
	GetUnitsResponse struct {
		Units []domain.Unit `json:"units"`
		Err   error         `json:"error"`
	}
	UpdateUnitRequest struct {
		Unit domain.Unit `json:"unit"`
	}
	UpdateUnitResponse struct {
		Unit *domain.Unit `json:"unit"`
		Err  error        `json:"error"`
	}
	DeleteUnitRequest struct {
		Id int64 `json:"id"`
	}
	DeleteUnitResponse struct {
		Err error `json:"error"`
	}
)
