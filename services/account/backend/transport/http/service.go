package http

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/account/backend/transport"
	"github.com/joyzem/proxy-project/services/account/dto"
	"github.com/joyzem/proxy-project/services/base"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)

	options = append(options, errorEncoder)
	router.Methods("POST").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.CreateAccount,
			decodeCreateAccountRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.GetAccounts,
			decodeGetAccountsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdateAccount,
			decodeUpdateAccountRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/accounts").Handler(
		kithttp.NewServer(
			svcEndpoints.DeleteAccount,
			decodeDeleteAccountRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.CreateAccountRequest
	return base.DecodeBody(r, &req)
}

func decodeGetAccountsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return dto.GetAccountsRequest{}, nil
}

func decodeDeleteAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.DeleteAccountRequest
	return base.DecodeBody(r, &req)
}

func decodeUpdateAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.UpdateAccountRequest
	return base.DecodeBody(r, &req)
}
