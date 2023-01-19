package transport

import "github.com/joyzem/proxy-project/services/product/domain"

type GetProductsResponse struct {
	Products []domain.Product `json:"products"`
	Err      error            `json:"error"`
}

type DeleteProductRequest struct {
	Id int `json:"id"`
}

type DeleteProductResponse struct {
	Err error `json:"error"`
}

type CreateProductRequest struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	UnitId int    `json:"unit_id"`
}

type CreateProductResponse struct {
	Product *domain.Product `json:"product"`
	Err     error           `json:"error"`
}

type UpdateProductTemplate struct {
	Product *domain.Product
	Units   []domain.Unit
}

type UpdateProductRequest struct {
	Product *domain.Product `json:"product"`
}
