package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/proxy-project/services/product/backend/service"
	"github.com/joyzem/proxy-project/services/product/dto"
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
		req := request.(dto.CreateProductRequest)
		product, err := s.CreateProduct(req.Name, req.Price, req.UnitId)
		return dto.CreateProductResponse{Product: product}, err
	}
}

func makeGetProductsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetProducts()
		return dto.GetProductsResponse{Products: products}, err
	}
}

func makeUpdateProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateProductRequest)
		service, err := s.UpdateProduct(req.Product)
		return dto.UpdateProductResponse{Product: service}, err
	}
}

func makeDeleteProductEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteProductRequest)
		err := s.DeleteProduct(req.Id)
		return dto.DeleteProductResponse{}, err
	}
}

func makeDeleteUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteUnitRequest)
		err := s.DeleteUnit(req.Id)
		return dto.DeleteUnitResponse{}, err
	}
}

func makeUpdateUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateUnitRequest)
		unit, err := s.UpdateUnit(req.Unit)
		return dto.UpdateUnitResponse{Unit: unit}, err
	}
}

func makeGetUnitsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		units, err := s.GetUnits()
		return dto.GetUnitsResponse{
			Units: units,
		}, err

	}
}

func makeCreateUnitEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateUnitRequest)
		unit, err := s.CreateUnit(req.Unit)
		return dto.CreateUnitResponse{Unit: unit}, err
	}
}
