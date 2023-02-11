package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joyzem/proxy-project/services/account/backend/implementation"
	"github.com/joyzem/proxy-project/services/account/backend/transport"
	"github.com/joyzem/proxy-project/services/base"

	kithttp "github.com/go-kit/kit/transport/http"
	httptransport "github.com/joyzem/proxy-project/services/account/backend/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	if err := implementation.InitDatabase(*db); err != nil {
		base.LogError(err)
	}

	repo := implementation.NewRepo(db)
	svc := implementation.NewService(repo)
	endpoints := transport.MakeEndpoints(svc)

	handler := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7073...")
	if err := http.ListenAndServe(":7073", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
