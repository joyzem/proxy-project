package utils

import (
	"fmt"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/dto"
	"github.com/levigross/grequests"
)

const (
	BACKEND_ADDRESS = "PRODUCT_BACKEND_ADDRESS"
)

// Получение адреса бэкэнда
func GetBackendAddress() string {
	// получение адреса из переменной окружения или использование http://localhost:7071
	address := base.GetEnv(BACKEND_ADDRESS, "http://localhost:7071")
	return address
}

// Получение единиц измерения с бэкэнда
func GetUnitsFromBackend() (*dto.GetUnitsResponse, error) {
	url := fmt.Sprintf("%s/units", GetBackendAddress())
	// отправка запроса и получение ответа
	resp, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	// парсинг ответа
	var data dto.GetUnitsResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Получение товаров с бэкэнда
func GetProductsFromBackend() (*dto.GetProductsResponse, error) {
	productsUrl := fmt.Sprintf("%s/products", GetBackendAddress())
	// отправка запроса и получение ответа
	resp, err := grequests.Get(productsUrl, nil)
	if err != nil {
		return nil, err
	}
	// парсинг ответа
	var data dto.GetProductsResponse
	if err := resp.JSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
