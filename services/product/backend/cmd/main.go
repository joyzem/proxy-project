package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/backend/implementation"
	"github.com/joyzem/proxy-project/services/product/backend/transport"
	httptransport "github.com/joyzem/proxy-project/services/product/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	// Подключение к БД
	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	// Создать таблицы, если не существуют
	if err := implementation.InitDatabase(*db); err != nil {
		base.LogError(errors.New(err.Error() + ": failed init"))
	}

	defer db.Close()

	// Создание сервиса
	productRepo := implementation.NewProductRepo(db)
	unitRepo := implementation.NewUnitRepository(db)
	svc := implementation.NewService(productRepo, unitRepo)

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7071...")
	if err := http.ListenAndServe(":7071", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

}
