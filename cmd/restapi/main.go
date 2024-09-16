package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/middleware"
)

func main() {
	appConfig, err := appconfig.LoadFromPath(context.Background(), "pkl/config/env.pkl")
	if err != nil {
		panic(err)
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(
		stopCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// # Load All Depedency Needed
	appDependency := bootstrap.Load(
		context.Background(),
		config.NewMySQLX(appConfig),
		config.NewNatsClient(appConfig),
		// add more depedency on here...
	)

	// # Init Go-Chi Router
	router := chi.NewRouter()

	// # Assign Global Middleware
	middleware.Global(appConfig, appDependency, router)

	// # Load HTTP Server API Configuration
	httpServerOption := config.NewHttpServerOption()
	httoServer, err := config.NewHTTPServer(

		// # Required
		appConfig,
		appDependency,
		router,

		// # Options
		httpServerOption.WithIdleTimeout(30*time.Second),
		httpServerOption.WithReadHeaderTimeout(5*time.Second),
		httpServerOption.WithReadTimeout(15*time.Second),
		httpServerOption.WithWriteTimeout(20*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to Load Config HTTP Server, got error %+v\n", err)
	}

	srv, err := httoServer.Create(context.Background())
	if err != nil {
		log.Fatalf("Failed to Load Initiator HTTP Server, got error %+v\n", err)
	}

	go func() {
		log.Printf("Start Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttpPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Service HTTP API on address %s:%d, listen and serve : %+v\n", appConfig.AppHost, appConfig.AppHttpPort, err)
		}
		log.Printf("Stop Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttpPort)
	}()

	<-stopCh

	srv.RegisterOnShutdown(func() {
		// Gracefully Stop Service and Close connection
		if appDependency.MySqlDB.Loaded() {
			appDependency.MySqlDB.Value().DB.Close()
		}

		if appDependency.NatsConn.Loaded() {
			appDependency.NatsConn.Value().Close()
		}

		log.Printf("Successfully Stop Service HTTP API on address %s:%d is exited properly\n", appConfig.AppHost, appConfig.AppHttpPort)
	})

	log.Printf("Perform shutdown with a maximum timeout of 30 seconds, Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttpPort)
	releaseCtx, releaseFn := context.WithTimeout(context.Background(), 30*time.Second)

	defer releaseFn()

	if err := srv.Shutdown(releaseCtx); err != nil {
		log.Printf("Shutdown Service HTTP API on address %s:%d, err : %+v\n", appConfig.AppHost, appConfig.AppHttpPort, err)
	}
}
