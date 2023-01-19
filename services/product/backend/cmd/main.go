package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/joyzem/proxy-project/services/product/backend/implementation"
	"github.com/joyzem/proxy-project/services/product/backend/service"
	"github.com/joyzem/proxy-project/services/product/backend/transport"
	httptransport "github.com/joyzem/proxy-project/services/product/backend/transport/http"
	"github.com/joyzem/proxy-project/services/utils"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	// Подключение к БД
	db, err := connectToDb()
	if err != nil {
		utils.LogError(err)
		os.Exit(-1)
	}

	// Создать таблицы, если не существуют
	if err := implementation.InitDatabase(*db); err != nil {
		utils.LogError(errors.New(err.Error() + ": failed init"))
	}

	defer db.Close()

	// Создание сервиса
	var svc service.Service
	{
		productRepo := implementation.NewProductRepo(db)
		unitRepo := implementation.NewUnitRepository(db)
		svc = implementation.NewService(productRepo, unitRepo)
	}

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	var h http.Handler
	{
		ocTracing := kitoc.HTTPServerTrace()
		serverOptions := []kithttp.ServerOption{ocTracing}
		h = httptransport.NewService(endpoints, serverOptions)
	}

	fmt.Println("Listening on 8080...")
	if err := http.ListenAndServe(":8080", h); err != nil {
		utils.LogError(err)
		os.Exit(-1)
	}

}

func connectToDb() (*sql.DB, error) {
	databaseHost := utils.GetEnv("DATABASE_HOST", "localhost")
	databaseUser := utils.GetEnv("DATABASE_USER", "rodion")
	connection := fmt.Sprintf("postgresql://%s:qwerty@%s:5432/proxy_db?sslmode=disable", databaseUser, databaseHost)
	var db *sql.DB
	// Проверка на доступность подключения к бд
	connectionAttempt := 0
	attemptsLimit := 5
	for ; connectionAttempt < attemptsLimit; connectionAttempt++ {
		var err error
		db, err = sql.Open("postgres", connection)
		if err != nil {
			utils.LogError(err)
			os.Exit(-1)
		}
		if err := db.Ping(); err != nil {
			utils.LogError(errors.New(
				"database is not responding; attempts: " + strconv.Itoa(connectionAttempt+1) + "/" + strconv.Itoa(attemptsLimit)))
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	if connectionAttempt == 5 {
		return nil, errors.New("failed to connect to database")
	}
	return db, nil
}
