package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/organization/backend/implementation"
	"github.com/joyzem/proxy-project/services/organization/backend/transport"
	httptransport "github.com/joyzem/proxy-project/services/organization/backend/transport/http"

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
		base.LogError(err)
	}

	defer db.Close()

	// Репозиторий товаров
	organizationRepo := implementation.NewOrganizationRepo(db)
	// Репозиторий единиц измерения
	// Создание сервиса
	svc := implementation.NewService(organizationRepo)

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	// Создание маршрутизатора
	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7072...")
	if err := http.ListenAndServe(":7072", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
