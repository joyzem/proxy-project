package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	sharedUtils "github.com/joyzem/proxy-project/services/utils"

	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/product/domain"
	"github.com/joyzem/proxy-project/services/product/frontend/transport"
	"github.com/joyzem/proxy-project/services/product/frontend/utils"
	"github.com/levigross/grequests"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := utils.GetProductsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	productPage, _ := template.ParseFiles("../static/html/products.html")
	productPage.Execute(w, products)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Redirect(w, r, "/product/products/", http.StatusBadRequest)
		return
	}
	body := transport.DeleteProductRequest{Id: id}
	options := utils.CreateJsonRequestOption(body)
	productsUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	resp, err := grequests.Delete(productsUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var deleteResponse transport.DeleteProductResponse
	err = json.Unmarshal(resp.Bytes(), &deleteResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/products/", http.StatusSeeOther)
}

func CreateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	units, err := utils.GetUnitsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sort.Slice(units, func(i, j int) bool {
		return units[i].Id < units[j].Id
	})
	createProductPage, err := template.ParseFiles("../static/html/create-product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = createProductPage.Execute(w, units)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	productName := r.FormValue("name")
	productPrice, priceError := strconv.Atoi(r.FormValue("price"))
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))
	if len(productName) == 0 || priceError != nil || unitErr != nil {
		http.Error(w, errors.New(sharedUtils.FIELDS_VALIDATION_ERROR).Error(), http.StatusUnprocessableEntity)
		return
	}
	request := transport.CreateProductRequest{
		Name:   productName,
		Price:  productPrice,
		UnitId: unitId,
	}
	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	options := utils.CreateJsonRequestOption(request)
	resp, err := grequests.Post(productUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data transport.CreateProductResponse
	err = json.Unmarshal(resp.Bytes(), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if data.Err == nil {
		http.Redirect(w, r, "/product/products/", http.StatusSeeOther)
	} else {
		http.Error(w, data.Err.Error(), http.StatusInternalServerError)
	}
}

func UpdateProductGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Redirect(w, r, "/product/products/", http.StatusBadRequest)
		return
	}
	products, err := utils.GetProductsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	units, err := utils.GetUnitsFromBackend()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var requestedProduct *domain.Product
	for i, product := range products {
		if products[i].Id == id {
			requestedProduct = &product
		}
	}
	if requestedProduct == nil {
		http.Error(w, "the product does not exist", http.StatusBadRequest)
		return
	}
	updateProductPage, err := template.ParseFiles("../static/html/update-product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := transport.UpdateProductTemplate{
		Product: requestedProduct,
		Units:   units,
	}
	updateProductPage.Execute(w, data)

}

func UpdateProductPostHandler(w http.ResponseWriter, r *http.Request) {
	product, err := getProductFromTheForm(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	productUrl := fmt.Sprintf("%s/products", utils.GetBackendAddress())
	request := transport.UpdateProductRequest{
		Product: product,
	}
	options := &grequests.RequestOptions{
		JSON: request,
	}
	_, err = grequests.Put(productUrl, options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/product/products/", http.StatusSeeOther)
}

func getProductFromTheForm(r *http.Request) (*domain.Product, error) {
	productId, _ := strconv.Atoi(r.FormValue("id"))
	productName := r.FormValue("name")
	productPrice, priceError := strconv.Atoi(r.FormValue("price"))
	unitId, unitErr := strconv.Atoi(r.FormValue("unit_id"))
	if len(productName) == 0 || priceError != nil || unitErr != nil {
		return nil, errors.New(sharedUtils.FIELDS_VALIDATION_ERROR)
	}
	return &domain.Product{
		Id:    productId,
		Name:  productName,
		Price: productPrice,
		Unit: domain.Unit{
			Id: unitId,
		},
	}, nil
}
