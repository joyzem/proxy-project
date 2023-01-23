package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/proxy-project/services/home/handlers"
	homeUtils "github.com/joyzem/proxy-project/services/home/utils"
	"github.com/joyzem/proxy-project/services/utils"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	productAddress := utils.GetEnv("PRODUCT_ADDRESS", "http://localhost:8081")
	homeUtils.HandleWithProxy(router, "/product/", productAddress)

	organizationAddress := utils.GetEnv("ORGANIZATION_ADDRESS", "http://localhost:8082")
	homeUtils.HandleWithProxy(router, "/organization/", organizationAddress)

	accountAddress := utils.GetEnv("ACCOUNT_ADDRESS", "http://localhost:8083")
	homeUtils.HandleWithProxy(router, "/account/", accountAddress)

	employeeAddress := utils.GetEnv("EMPLOYEE_ADDRESS", "http://localhost:8084")
	homeUtils.HandleWithProxy(router, "/employee/", employeeAddress)

	customerAddress := utils.GetEnv("CUSTOMER_ADDRESS", "http://localhost:8085")
	homeUtils.HandleWithProxy(router, "/customer/", customerAddress)

	proxyAddress := utils.GetEnv("PROXY_ADDRESS", "http://localhost:8086")
	homeUtils.HandleWithProxy(router, "/proxy/", proxyAddress)

	router.HandleFunc("/", handlers.IndexHandler)

	router.HandleFunc("/about-golang/", handlers.GolangHandler)

	router.HandleFunc("/lab-{id:[0-8]}/", handlers.LabsHandler)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	router.PathPrefix("/home/static/").Handler(http.StripPrefix("/home/static/", http.FileServer(http.Dir("../static"))))

	return router
}
