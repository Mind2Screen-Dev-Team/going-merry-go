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

	repo_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/impl"
	service_impl "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/impl"
)

func main() {
	// # Load App Config
	appConfig, err := appconfig.LoadFromPath(context.Background(), "pkl/config/env.pkl")
	if err != nil {
		panic(err)
	}

	// # Handle Gracefully Shutdown Signal
	stopCh := make(chan os.Signal, 1)
	signal.Notify(
		stopCh,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	appDependency, appService := LoadApplication(appConfig)
	_ = appService

	// # Init Go-Chi Router
	router := chi.NewRouter()

	// # Assign Global Middleware
	middleware.Global(appConfig, router)

	// # Load HTTP Server API Configuration
	httpServerOption := config.NewHttpServerOption()
	httoServer, err := config.NewHTTPServer(

		// # Required
		appConfig,
		appDependency,
		router,

		// # Options
		httpServerOption.WithIdleTimeout(
			time.Duration(appConfig.AppHttp.IdleTimeout)*time.Second,
		),
		httpServerOption.WithReadHeaderTimeout(
			time.Duration(appConfig.AppHttp.ReadHeaderTimeout)*time.Second,
		),
		httpServerOption.WithReadTimeout(
			time.Duration(appConfig.AppHttp.ReadTimeout)*time.Second,
		),
		httpServerOption.WithWriteTimeout(
			time.Duration(appConfig.AppHttp.WriteTimeout)*time.Second,
		),
	)
	if err != nil {
		log.Fatalf("Failed to Load Config HTTP Server, got error %+v\n", err)
	}

	srv, err := httoServer.Create(context.Background())
	if err != nil {
		log.Fatalf("Failed to Load Initiator HTTP Server, got error %+v\n", err)
	}

	go func() {
		log.Printf("Start Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttp.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Service HTTP API on address %s:%d, listen and serve : %+v\n", appConfig.AppHost, appConfig.AppHttp.Port, err)
		}
		log.Printf("Stop Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttp.Port)
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

		log.Printf("Successfully Stop Service HTTP API on address %s:%d is exited properly\n", appConfig.AppHost, appConfig.AppHttp.Port)
	})

	log.Printf("Perform shutdown with a maximum timeout of 30 seconds, Service HTTP API on address %s:%d\n", appConfig.AppHost, appConfig.AppHttp.Port)
	releaseCtx, releaseFn := context.WithTimeout(context.Background(), 30*time.Second)

	defer releaseFn()

	if err := srv.Shutdown(releaseCtx); err != nil {
		log.Printf("Shutdown Service HTTP API on address %s:%d, err : %+v\n", appConfig.AppHost, appConfig.AppHttp.Port, err)
	}
}

func LoadApplication(appConfig *appconfig.AppConfig) (dep *bootstrap.AppDependency, service *bootstrap.AppService) {
	var repo *bootstrap.AppRepository

	// # Load All Dependency
	dep = bootstrap.LoadDependency(
		context.Background(),
		config.NewMySQLX(appConfig),
		config.NewNatsClient(appConfig),

		// add more on here...
	)

	// # Load All Repository
	repo = bootstrap.LoadRepository(
		context.Background(),

		// Link Dependency
		dep,

		// # List Of Repository
		repo_impl.NewUserRepoImpl(),

		// add more on here...
	)

	// # Load All Service
	service = bootstrap.LoadService(
		context.Background(),
		// Link Dependency and Service
		dep,
		repo,

		// # List Of Service
		service_impl.NewUserServiceImpl(),

		// add more on here...
	)

	return
}
