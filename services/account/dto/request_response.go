package dto

import "github.com/joyzem/proxy-project/services/account/domain"

type (
	CreateAccountRequest struct {
		BankName           string `json:"bank_name"`
		BankIdentityNumber string `json:"bank_identity_number"`
	}
	CreateAccountResponse struct {
		Account *domain.Account `json:"account,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	GetAccountsRequest struct {
	}
	GetAccountsResponse struct {
		Accounts []domain.Account `json:"accounts,omitempty"`
		Err      string           `json:"error,omitempty"`
	}
	UpdateAccountRequest struct {
		Account domain.Account `json:"account"`
	}
	UpdateAccountResponse struct {
		Account *domain.Account `json:"account,omitempty"`
		Err     string          `json:"error,omitempty"`
	}
	DeleteAccountRequest struct {
		Id int `json:"id"`
	}
	DeleteAccountResponse struct {
		Err string `json:"error,omitempty"`
	}
)
