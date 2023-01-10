package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	product "github.com/joyzem/proxy-project/services/product/backend"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
	CreateUnit    endpoint.Endpoint
	GetUnits      endpoint.Endpoint
	UpdateUnit    endpoint.Endpoint
	DeleteUnit    endpoint.Endpoint
}

func MakeEndpoints(s product.Service) Endpoints {
	return Endpoints{
		CreateProduct: makeCreateProductEndpoint(s),
		GetProducts:   makeGetProductsEndpoint(s),
		UpdateProduct: makeUpdateProductEndpoint(s),
		DeleteProduct: makeDeleteProductEndpoint(s),
		CreateUnit:    makeCreateUnitEndpoint(s),
		GetUnits:      makeGetUnitsEndpoint(s),
		UpdateUnit:    makeUpdateUnitEndpoint(s),
		DeleteUnit:    makeDeleteUnitEndpoint(s),
	}
}

func makeCreateProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		product, err := s.CreateProduct(ctx, req.Name, req.Price, req.Unit)
		return CreateProductResponse{Product: product, Err: err}, err
	}
}

func makeGetProductsEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetProducts(ctx)
		return GetProductsResponse{
			Products: products,
			Err:      err,
		}, err

	}
}

func makeUpdateProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		product, err := s.UpdateProduct(ctx, req.Product)
		return UpdateProductResponse{Product: product, Err: err}, err
	}
}

func makeDeleteProductEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := s.DeleteProduct(ctx, req.Id)
		return DeleteProductResponse{Err: err}, err
	}
}

func makeDeleteUnitEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUnitRequest)
		err := s.DeleteUnit(ctx, req.Id)
		return DeleteUnitResponse{Err: err}, err
	}
}

func makeUpdateUnitEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUnitRequest)
		unit, err := s.UpdateUnit(ctx, req.Unit)
		return UpdateUnitResponse{Unit: unit, Err: err}, err
	}
}

func makeGetUnitsEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		units, err := s.GetUnits(ctx)
		return GetUnitsResponse{
			Units: units,
			Err:   err,
		}, nil

	}
}

func makeCreateUnitEndpoint(s product.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUnitRequest)
		unit, err := s.CreateUnit(ctx, req.Unit)
		return CreateUnitResponse{Unit: unit, Err: err}, err
	}
}
