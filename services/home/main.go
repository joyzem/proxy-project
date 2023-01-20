package main

import (
	"fmt"
	"net/http"

	"github.com/joyzem/proxy-project/services/utils"
)

func main() {

	productAddress := utils.GetEnv("PRODUCT_ADDRESS", "http://localhost:8081")
	utils.HandleWithProxy("/product/", productAddress)
	organizationAddress := utils.GetEnv("ORGANIZATION_ADDRESS", "http://localhost:8082")
	utils.HandleWithProxy("/organization/", organizationAddress)
	accountAddress := utils.GetEnv("ACCOUNT_ADDRESS", "http://localhost:8083")
	utils.HandleWithProxy("/account/", accountAddress)
	employeeAddress := utils.GetEnv("EMPLOYEE_ADDRESS", "http://localhost:8084")
	utils.HandleWithProxy("/employee/", employeeAddress)
	customerAddress := utils.GetEnv("CUSTOMER_ADDRESS", "http://localhost:8085")
	utils.HandleWithProxy("/customer/", customerAddress)
	proxyAddress := utils.GetEnv("PROXY_ADDRESS", "http://localhost:8086")
	utils.HandleWithProxy("/proxy/", proxyAddress)

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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../static"))))

	fmt.Println("Listening on 80...")
	http.ListenAndServe(":80", nil)
}
