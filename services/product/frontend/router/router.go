package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/product/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/html/index.html")
	})

	router.HandleFunc("/product/products/", handlers.ProductsHandler)
	router.HandleFunc("/product/products/delete/", handlers.DeleteProductHandler).Methods(http.MethodPost)

	router.HandleFunc("/product/products/create/", handlers.CreateProductGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/products/create/", handlers.CreateProductPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/product/products/update/{id}", handlers.UpdateProductGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/products/update/", handlers.UpdateProductPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/product/units/", handlers.UnitsHandler)
	router.HandleFunc("/product/units/delete/", handlers.DeleteUnitHandler).Methods(http.MethodPost)

	router.HandleFunc("/product/units/create/", handlers.CreateUnitGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/units/create/", handlers.CreateUnitPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/product/units/update/{id}", handlers.UpdateUnitGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/units/update/", handlers.UpdateUnitPostHandler).Methods(http.MethodPost)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/product/static/").Handler(http.StripPrefix("/product/static/", http.FileServer(http.Dir("../static"))))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
	})
	handler := c.Handler(router)
	return handler
}
