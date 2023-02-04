package utils

import (
	"encoding/json"
	"fmt"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/dto"
	"github.com/levigross/grequests"
)

const (
	BACKEND_ADDRESS = "PRODUCT_BACKEND_ADDRESS"
)

func GetBackendAddress() string {
	address := base.GetEnv(BACKEND_ADDRESS, "http://localhost:7071")
	return address
}

func GetUnitsFromBackend() (*dto.GetUnitsResponse, error) {
	url := fmt.Sprintf("%s/units", GetBackendAddress())
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data dto.GetUnitsResponse
	if err := json.Unmarshal(resp.Bytes(), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func GetProductsFromBackend() (*dto.GetProductsResponse, error) {
	productsUrl := fmt.Sprintf("%s/products", GetBackendAddress())
	resp, err := grequests.Get(productsUrl, nil)
	if err != nil {
		return nil, err
	}
	var data dto.GetProductsResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
