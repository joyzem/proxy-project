package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	product "github.com/joyzem/proxy-project/services/product/backend"
	"github.com/joyzem/proxy-project/services/product/backend/implementation"
	"github.com/joyzem/proxy-project/services/product/backend/transport"
	httptransport "github.com/joyzem/proxy-project/services/product/backend/transport/http"
	"github.com/joyzem/proxy-project/services/utils"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

	logger := newLogger()

	// Подключение к БД
	db, err := connectToDb(&logger)
	if err != nil {
		utils.LogError(&logger, err)
		os.Exit(-1)
	}

	// Создать таблицы, если не существуют
	if err := implementation.InitDatabase(*db); err != nil {
		utils.LogError(&logger, errors.New(err.Error()+": failed init"))
	}

	defer db.Close()
	defer level.Info(logger).Log("msg", "service ended")

	// Создание сервиса
	var svc product.Service
	{
		productRepo, err := implementation.NewProductRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		unitRepo, err := implementation.NewUnitRepository(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = implementation.NewService(productRepo, unitRepo, logger)
	}

	// Создание эндпоинтов
	endpoints := transport.MakeEndpoints(svc)

	var h http.Handler
	{
		ocTracing := kitoc.HTTPServerTrace()
		serverOptions := []kithttp.ServerOption{ocTracing}
		h = httptransport.NewService(endpoints, serverOptions, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}

func newLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	return logger
}

func connectToDb(logger *log.Logger) (*sql.DB, error) {
	connection := "postgresql://user:qwerty@database:5432/proxy_db?sslmode=disable"
	var db *sql.DB
	// Проверка на доступность подключения к бд
	connectionAttempt := 0
	attemptsLimit := 5
	for ; connectionAttempt < attemptsLimit; connectionAttempt++ {
		var err error
		db, err = sql.Open("postgres", connection)
		if err != nil {
			utils.LogError(logger, err)
			os.Exit(-1)
		}
		if err := db.Ping(); err != nil {
			utils.LogError(logger, errors.New(
				"database is not responding; attempts: "+strconv.Itoa(connectionAttempt+1)+"/"+strconv.Itoa(attemptsLimit)))
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
