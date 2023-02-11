package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/employee/backend/implementation"
	"github.com/joyzem/proxy-project/services/employee/backend/transport"
	httptransport "github.com/joyzem/proxy-project/services/employee/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	if err := implementation.InitDatabase(db); err != nil {
		base.LogError(err)
	}

	defer db.Close()

	employeeRepo := implementation.NewEmployeeRepo(db)
	svc := implementation.NewEmployeeService(employeeRepo)

	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7074...")
	if err := http.ListenAndServe(":7074", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
