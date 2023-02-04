package dto

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
		Product *domain.Product `json:"product,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	GetProductsRequest struct {
	}
	GetProductsResponse struct {
		Products []domain.Product `json:"products,omitempty"`
		Err      string           `json:"error,omitempty"`
	}
	UpdateProductRequest struct {
		Product domain.Product `json:"product"`
	}
	UpdateProductResponse struct {
		Product *domain.Product `json:"product,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	DeleteProductRequest struct {
		Id int `json:"id"`
	}
	DeleteProductResponse struct {
		Err string `json:"error,omitempty"`
	}
)
