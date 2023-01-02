package transport

import "github.com/joyzem/proxy-project/services/product"

type (
	CreateProductRequest struct {
		Product product.Product `json:"product"`
	}
	CreateProductResponse struct {
		Err error `json:"error"`
	}
	GetProductsRequest struct {
	}
	GetProductsResponse struct {
		Products []product.Product `json:"products"`
		Err      error             `json:"error"`
	}
	GetProductRequest struct {
		Id int `json:"id"`
	}
	GetProductResponse struct {
		Product product.Product `json:"product"`
		Err     error           `json:"error"`
	}
	UpdateProductRequest struct {
		Id      int             `json:"id"`
		Product product.Product `json:"product"`
	}
	UpdateProductResponse struct {
		Err error `json:"error"`
	}
	DeleteProductRequest struct {
		Id int `json:"id"`
	}
	DeleteProductResponse struct {
		Err error `json:"error"`
	}
)
