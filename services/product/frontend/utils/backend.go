package utils

import (
	"encoding/json"
	"fmt"

	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/product/frontend/transport"
	"github.com/joyzem/proxy-project/services/utils"
	"github.com/levigross/grequests"
)

const (
	BACKEND_ADDRESS = "PRODUCT_BACKEND_ADDRESS"
)

func GetBackendAddress() string {
	address := utils.GetEnv(BACKEND_ADDRESS, "http://localhost:8080")
	return address
}

func GetUnitsFromBackend() ([]domain.Unit, error) {
	url := fmt.Sprintf("%s/units", GetBackendAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data transport.GetUnitsResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		return nil, err
	}
	return data.Units, nil
}

func GetProductsFromBackend() ([]domain.Product, error) {
	productsUrl := fmt.Sprintf("%s/products", GetBackendAddress())
	resp, err := grequests.Get(productsUrl, nil)
	if err != nil {
		return nil, err
	}
	var data transport.GetProductsResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		return nil, err
	}
	return data.Products, nil
}
