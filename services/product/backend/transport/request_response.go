package transport

import product "github.com/joyzem/proxy-project/services/product/backend"

type (
	CreateProductRequest struct {
		Name  string       `json:"name"`
		Price int32        `json:"price"`
		Unit  product.Unit `json:"unit"`
	}
	CreateProductResponse struct {
		Product *product.Product `json:"product"`
		Err     error            `json:"error"`
	}
	GetProductsRequest struct {
	}
	GetProductsResponse struct {
		Products []product.Product `json:"products"`
		Err      error             `json:"error"`
	}
	UpdateProductRequest struct {
		Product product.Product `json:"product"`
	}
	UpdateProductResponse struct {
		Product *product.Product `json:"product"`
		Err     error            `json:"error"`
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
		Unit *product.Unit `json:"unit"`
		Err  error         `json:"error"`
	}
	GetUnitsRequest struct {
	}
	GetUnitsResponse struct {
		Units []product.Unit `json:"units"`
		Err   error          `json:"error"`
	}
	UpdateUnitRequest struct {
		Unit product.Unit `json:"unit"`
	}
	UpdateUnitResponse struct {
		Unit *product.Unit `json:"unit"`
		Err  error         `json:"error"`
	}
	DeleteUnitRequest struct {
		Id int64 `json:"id"`
	}
	DeleteUnitResponse struct {
		Err error `json:"error"`
	}
)
