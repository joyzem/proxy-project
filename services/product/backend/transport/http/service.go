package http

import (
	"context"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/product/backend/transport"
	"github.com/joyzem/proxy-project/services/utils"
)

var (
	ErrBadRouting = errors.New("bad Routing")
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	var (
		router       = mux.NewRouter()
		errorEncoder = kithttp.ServerErrorEncoder(utils.EncodeErrorResponse)
	)

	options = append(options, errorEncoder)

	router.Methods("POST").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProduct,
			decodeCreateProductRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProducts,
			decodeGetProductsRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateProduct,
			decodeUpdateProductRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProduct,
			decodeDeleteProductRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("POST").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateUnit,
			decodeCreateUnitRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.GetUnits,
			decodeGetUnitsRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateUnit,
			decodeUpdateUnitRequest,
			utils.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/units").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteUnit,
			decodeDeleteUnitRequest,
			utils.EncodeResponse,
			options...,
		))
	return router
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateProductRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.GetProductsRequest
	return req, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.UpdateProductRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.DeleteProductRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}

func decodeCreateUnitRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateUnitRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}

func decodeGetUnitsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.GetUnitsRequest
	return req, nil
}

func decodeUpdateUnitRequest(_ context.Context, r *http.Request) (reques interface{}, err error) {
	var req transport.UpdateUnitRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}

func decodeDeleteUnitRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.DeleteUnitRequest
	if err := utils.DecodeBody(r, &req); err != nil {
		utils.LogError(err)
		return nil, err
	}
	return req, nil
}
