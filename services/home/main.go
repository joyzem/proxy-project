package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	handleProxy("/product/", ":8081")
	handleProxy("/organization/", ":8082")
	handleProxy("/account/", ":8083")
	handleProxy("/employee/", ":8084")
	handleProxy("/customer/", ":8085")
	handleProxy("/proxy/", ":8086")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/about-proxy/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/proxy.html")
	})
	http.HandleFunc("/about-account/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/account.html")
	})
	http.HandleFunc("/about-customer/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/customer.html")
	})
	http.HandleFunc("/about-employee/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/employee.html")
	})
	http.HandleFunc("/about-organization/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/organization.html")
	})
	http.HandleFunc("/about-product/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/product.html")
	})

	http.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir("../../share"))))
	http.Handle("/home/js/", http.StripPrefix("/home/js/", http.FileServer(http.Dir("views"))))
	fmt.Println("Home service started...")

	http.ListenAndServe(":80", nil)
}

// Обработать прокси запрос
// path - /path/to/endpoint
// port - :xxxx
func handleProxy(path string, port string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse("http://localhost" + port)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})
}
