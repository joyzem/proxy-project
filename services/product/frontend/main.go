package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/product/products/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/products.html")
	})
	http.HandleFunc("/product/units/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/units.html")
	})
	http.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir("../../../share"))))
	http.ListenAndServe(":8081", nil)
	fmt.Println("Products service started...")
}
