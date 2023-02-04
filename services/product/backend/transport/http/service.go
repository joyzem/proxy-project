package http

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/backend/transport"
	"github.com/joyzem/proxy-project/services/product/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {

	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)

	options = append(options, errorEncoder)

	router.Methods("POST").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProduct,
			decodeCreateProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProducts,
			decodeGetProductsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateProduct,
			decodeUpdateProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProduct,
			decodeDeleteProductRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("POST").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateUnit,
			decodeCreateUnitRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.GetUnits,
			decodeGetUnitsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateUnit,
			decodeUpdateUnitRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteUnit,
			decodeDeleteUnitRequest,
			base.EncodeResponse,
			options...,
		))
	return router
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.CreateProductRequest
	return base.DecodeBody(r, &req)
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return dto.GetProductsRequest{}, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.UpdateProductRequest
	return base.DecodeBody(r, &req)
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.DeleteProductRequest
	return base.DecodeBody(r, &req)
}

func decodeCreateUnitRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.CreateUnitRequest
	return base.DecodeBody(r, &req)
}

func decodeGetUnitsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.GetUnitsRequest
	return req, nil
}

func decodeUpdateUnitRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.UpdateUnitRequest
	return base.DecodeBody(r, &req)
}

func decodeDeleteUnitRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.DeleteUnitRequest
	return base.DecodeBody(r, &req)
}
