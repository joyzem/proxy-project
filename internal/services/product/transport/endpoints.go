package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/proxy-project/internal/services/product"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	GetProduct    endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
}

func MakeEndpoints(s product.Service) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		GetProducts:   makeGetProductsEndpoint(s),
		GetProduct:    makeGetProductEndpoint(s),
		UpdateProduct: makeUpdateProductEndpoint(s),
		DeleteProduct: makeDeleteProductEndpoint(s),
	}
}

func makeCreateProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		err := s.CreateProduct(ctx, req.Product)
		return CreateProductResponse{Err: err}, nil
	}
}

func makeGetProductsEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetProducts(ctx)
		return GetProductsResponse{
			Products: products,
			Err:      err,
		}, nil

	}
}

func makeGetProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		product, err := s.GetProduct(ctx, req.Id)
		return GetProductResponse{Product: product, Err: err}, nil
	}
}

func makeUpdateProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		err := s.UpdateProduct(ctx, req.Id, req.Product)
		return UpdateProductResponse{Err: err}, nil
	}
}

func makeDeleteProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := s.DeleteProduct(ctx, req.Id)
		return DeleteProductResponse{Err: err}, nil
	}
}
