package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/product/transport"
)

var (
	ErrBadRouting = errors.New("Bad Routing")
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
	logger log.Logger,
) http.Handler {
	var (
		router       = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)

	options = append(options, errorLogger, errorEncoder)

	router.Methods("POST").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateProduct,
			decodeCreateProductRequest,
			encodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProducts,
			decodeGetProductsRequest,
			encodeResponse,
			options...,
		))

	router.Methods("GET").Path("/products/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.GetProduct,
			decodeGetProductRequest,
			encodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/products/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateProduct,
			decodeUpdateProductRequest,
			encodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/products/{id}").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteProduct,
			decodeDeleteProductRequest,
			encodeResponse,
			options...,
		))
	return router
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateProductRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Product); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.GetProductsRequest
	return req, nil
}

func decodeGetProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	intId, err := strconv.Atoi(id)
	if !ok || err != nil {
		return nil, ErrBadRouting
	}
	return transport.GetProductRequest{Id: intId}, nil
}

func decodeUpdateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	intId, err := strconv.Atoi(id)
	if !ok || err != nil {
		return nil, ErrBadRouting
	}
	return transport.UpdateProductRequest{Id: intId}, nil
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	intId, err := strconv.Atoi(id)
	if !ok || err != nil {
		return nil, ErrBadRouting
	}
	return transport.DeleteProductRequest{Id: intId}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	return http.StatusInternalServerError

}
