run-product:
	go run services/product/frontend/main.go &

run-account:
	go run services/account/frontend/main.go &

run-customer:
	go run services/customer/frontend/main.go &

run-employee:
	go run services/employee/frontend/main.go &

run-home:
	go run services/home/main.go &

run-organization:
	go run services/organization/frontend/main.go &

run-proxy:
	go run services/proxy/frontend/main.go &

run-all: run-product run-account run-customer run-employee run-home run-organization run-proxy

stop:
	pkill -f "product"
	pkill -f "account"
	pkill -f "customer"
	pkill -f "employee"
	pkill -f "home"
	pkill -f "organization"
	pkill -f "proxy"