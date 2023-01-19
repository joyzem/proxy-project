package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/proxy-project/services/product/backend/service"
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

func MakeEndpoints(s service.Service) Endpoints {
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

func makeCreateProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		service, err := s.CreateProduct(ctx, req.Name, req.Price, req.UnitId)
		return CreateProductResponse{Product: service, Err: err}, err
	}
}

func makeGetProductsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetProducts(ctx)
		return GetProductsResponse{
			Products: products,
			Err:      err,
		}, err

	}
}

func makeUpdateProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		service, err := s.UpdateProduct(ctx, req.Product)
		return UpdateProductResponse{Product: service, Err: err}, err
	}
}

func makeDeleteProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		err := s.DeleteProduct(ctx, req.Id)
		return DeleteProductResponse{Err: err}, err
	}
}

func makeDeleteUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUnitRequest)
		err := s.DeleteUnit(ctx, req.Id)
		return DeleteUnitResponse{Err: err}, err
	}
}

func makeUpdateUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUnitRequest)
		unit, err := s.UpdateUnit(ctx, req.Unit)
		return UpdateUnitResponse{Unit: unit, Err: err}, err
	}
}

func makeGetUnitsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		units, err := s.GetUnits(ctx)
		return GetUnitsResponse{
			Units: units,
			Err:   err,
		}, nil

	}
}

func makeCreateUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUnitRequest)
		unit, err := s.CreateUnit(ctx, req.Unit)
		return CreateUnitResponse{Unit: unit, Err: err}, err
	}
}
