package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/product/dto"
	"github.com/joyzem/proxy-project/services/product/frontend/utils"
	"github.com/levigross/grequests"
)

// Обработчик страницы всех товаров
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// получить товары с бэка
	response, _ := utils.GetProductsFromBackend()
	if response.Err != "" {
		http.Error(w, response.Err, http.StatusInternalServerError)
		return
	}
	// создать шаблон
	productPage, _ := template.ParseFiles("../static/html/products.html")
	productPage.Execute(w, response.Products)
}

// Обработчик удаления товара
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// прочитать id
	id, _ := strconv.Atoi(r.FormValue("id"))
	// адрес бэка
	productsUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	// отправить запрос на удаление и получить ответ
	resp, _ := grequests.Delete(productsUrl, &grequests.RequestOptions{
		JSON: dto.DeleteProductRequest{Id: id},
	})
	// распарсить ответ
	var deleteResponse dto.DeleteProductResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	// редирект на все товары
	http.Redirect(w, r, "/product/products", http.StatusSeeOther)
}

// Обработчик страницы добавления товара
func CreateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	// получить единицы измерения
	units, _ := utils.GetUnitsFromBackend()
	if units.Err != "" {
		http.Error(w, units.Err, http.StatusInternalServerError)
		return
	}
	// отсортировать по id
	sort.Slice(units.Units, func(i, j int) bool {
		return units.Units[i].Id < units.Units[j].Id
	})
	// шаблон добавления товара
	createProductPage, _ := template.ParseFiles("../static/html/create-product.html")
	createProductPage.Execute(w, units.Units)
}

// Обработчик запроса на добавление товара
func CreateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	// наименование товара
	productName := r.FormValue("name")
	// цена товара
	productPrice, priceError := strconv.Atoi(r.FormValue("price"))
	// id ед.изм.
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))
	// валидация форм
	if productName == "" || priceError != nil || unitErr != nil {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	// создание структуры запроса
	request := dto.CreateProductRequest{
		Name:   productName,
		Price:  productPrice,
		UnitId: unitId,
	}
	// адрес бэка
	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	// отправить запрос на добавление товара и получение ответа
	resp, _ := grequests.Post(productUrl, &grequests.RequestOptions{
		JSON: request,
	})
	// парсинг ответа
	var data dto.CreateProductResponse
	resp.JSON(&data)
	if data.Err != "" {
		http.Error(w, data.Err, http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/product/products", http.StatusSeeOther)
	}
}

// Обработчик страницы обновления продукта
func UpdateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// получение продуктов с бэка
	getProductsResponse, _ := utils.GetProductsFromBackend()
	if getProductsResponse.Err != "" {
		http.Error(w, getProductsResponse.Err, http.StatusInternalServerError)
		return
	}
	// поиск нужного продукта
	var requestedProduct *domain.Product
	for i, product := range getProductsResponse.Products {
		if getProductsResponse.Products[i].Id == id {
			requestedProduct = &product
		}
	}
	// продукт не найден
	if requestedProduct == nil {
		http.Error(w, "the product does not exist", http.StatusBadRequest)
		return
	}
	// получение единиц измерения с бэка
	unitsResp, _ := utils.GetUnitsFromBackend()
	if unitsResp.Err != "" {
		http.Error(w, unitsResp.Err, http.StatusInternalServerError)
		return
	}
	// структура данных для шаблона
	type UpdateProductTemplate struct {
		Product *domain.Product
		Units   []domain.Unit
	}
	data := UpdateProductTemplate{
		Product: requestedProduct,
		Units:   unitsResp.Units,
	}
	// шаблон страницы
	updateProductPage, _ := template.ParseFiles("../static/html/update-product.html")
	updateProductPage.Execute(w, data)

}

// Обработчик запроса на обновление товара
func UpdateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	// парсинг товара
	productId, _ := strconv.Atoi(r.FormValue("id"))
	productName := r.FormValue("name")
	productPrice, priceError := strconv.Atoi(r.FormValue("price"))
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))
	// валидация форм
	if len(productName) == 0 || priceError != nil || unitErr != nil {
		http.Error(w, base.FIELDS_VALIDATION_ERROR, http.StatusUnprocessableEntity)
		return
	}
	product := domain.Product{
		Id:    productId,
		Name:  productName,
		Price: productPrice,
		Unit: domain.Unit{
			Id: unitId,
		},
	}
	// адрес бэка
	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	// тело запроса
	request := dto.UpdateProductRequest{
		Product: product,
	}
	// отправка запроса на обновление и получение ответа
	resp, _ := grequests.Put(productUrl, &grequests.RequestOptions{
		JSON: request,
	})
	// парсинг ответа
	var updateResponse dto.UpdateProductResponse
	resp.JSON(&updateResponse)
	if updateResponse.Err != "" {
		http.Error(w, updateResponse.Err, http.StatusInternalServerError)
		return
	}
	// возврат на страницу продуктов
	http.Redirect(w, r, "/product/products", http.StatusSeeOther)
}
