package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/proxy-project/services/account/backend/service"
	"github.com/joyzem/proxy-project/services/account/dto"
)

type Endpoints struct {
	CreateAccount endpoint.Endpoint
	GetAccounts   endpoint.Endpoint
	UpdateAccount endpoint.Endpoint
	DeleteAccount endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		CreateAccount: makeCreateAccountEndpoint(s),
		GetAccounts:   makeGetAccountsEndpoint(s),
		UpdateAccount: makeUpdateAccountEndpoint(s),
		DeleteAccount: makeDeleteAccountEndpoint(s),
	}
}

func makeCreateAccountEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CreateAccountRequest)
		account, err := s.CreateAccount(req.BankName, req.BankIdentityNumber)
		return dto.CreateAccountResponse{Account: account}, err
	}
}

func makeGetAccountsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		accounts, err := s.GetAccounts()
		return dto.GetAccountsResponse{Accounts: accounts}, err
	}
}

func makeUpdateAccountEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UpdateAccountRequest)
		account, err := s.UpdateAccount(req.Account)
		return dto.UpdateAccountResponse{Account: account}, err
	}
}

func makeDeleteAccountEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteAccountRequest)
		err := s.DeleteAccount(req.Id)
		return dto.DeleteAccountResponse{}, err
	}
}
