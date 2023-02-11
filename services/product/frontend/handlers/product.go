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
	products, _ := utils.GetProductsFromBackend()
	if products.Err != "" {
		http.Error(w, products.Err, http.StatusInternalServerError)
		return
	}
	// получить единицы измерения с бэка
	units, _ := utils.GetUnitsFromBackend()
	if units.Err != "" {
		http.Error(w, units.Err, http.StatusInternalServerError)
		return
	}
	type ProductPageItemTemplate struct {
		Product domain.Product
		Unit    domain.Unit
	}
	templateItems := []ProductPageItemTemplate{}
	for _, product := range products.Products {
		templateItem := ProductPageItemTemplate{}
		templateItem.Product = product
		var productUnit domain.Unit
		for _, unit := range units.Units {
			if unit.Id == product.UnitId {
				productUnit = unit
				break
			}
		}
		templateItem.Unit = productUnit
		templateItems = append(templateItems, templateItem)
	}
	// создать шаблон
	productPage, _ := template.ParseFiles("../static/html/products.html")
	productPage.Execute(w, templateItems)
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
	// получение продукта с бэка
	url := fmt.Sprintf("%s/products/%d", utils.GetBackendAddress(), id)
	resp, _ := grequests.Get(url, &grequests.RequestOptions{
		JSON: dto.ProductByIdRequest{Id: id},
	})
	var product dto.ProductByIdResponse
	resp.JSON(&product)
	// продукт не найден
	if product.Err != "" {
		http.Error(w, product.Err, http.StatusBadRequest)
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
		Product: product.Product,
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
		Id:     productId,
		Name:   productName,
		Price:  productPrice,
		UnitId: unitId,
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
