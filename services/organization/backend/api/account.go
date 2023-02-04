package api

import (
	"encoding/json"
	"fmt"

	"github.com/joyzem/proxy-project/services/account/domain"
	"github.com/joyzem/proxy-project/services/account/dto"
	"github.com/joyzem/proxy-project/services/base"
	"github.com/levigross/grequests"
)

type AccountApi interface {
	GetAccount(id int) (*domain.Account, error)
}

type accountAPIClient struct {
	accountsUrl string
}

func (c *accountAPIClient) GetAccount(id int) (*domain.Account, error) {

	resp, err := grequests.Get(c.accountsUrl+"/accounts", nil)
	if err != nil {
		return nil, err
	}
	var accountDto dto.GetAccountsResponse
	if err := json.Unmarshal(resp.Bytes(), &accountDto); err != nil {
		return nil, err
	}
	if accountDto.Err != "" {
		return nil, err
	}
	fmt.Printf("accountDto: %v\n", accountDto)
	var requestedAccount *domain.Account
	for _, acc := range accountDto.Accounts {
		if acc.Id == id {
			requestedAccount = &acc
		}
	}
	return requestedAccount, nil
}

func NewAccountApiClient() AccountApi {
	return &accountAPIClient{accountsUrl: getAccountsUrl()}
}

func getAccountsUrl() string {
	return base.GetEnv("ACCOUNT_API_BASE_URL", "http://localhost:7073")
}
