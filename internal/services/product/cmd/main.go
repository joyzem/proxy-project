package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/joyzem/proxy-project/internal/services/product"
	"github.com/joyzem/proxy-project/internal/services/product/implementation"
	"github.com/joyzem/proxy-project/internal/services/product/transport"
	httptransport "github.com/joyzem/proxy-project/internal/services/product/transport/http"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

	// Подключение к БД
	var db *sql.DB
	{
		var err error
		connection := "user=rodion password=qwerty dbname=proxy_db sslmode=disable"
		db, err = sql.Open("postgres", connection)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	logger := newLogger()
	defer level.Info(logger).Log("msg", "service ended")

	// Создание сервиса
	var svc product.Service
	{
		repository, err := implementation.NewRepository(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = implementation.NewService(repository)
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
