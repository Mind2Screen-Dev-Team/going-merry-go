package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/middleware"
)

func main() {
	// # Parse App Config Path
	cfgPath := flag.String("cfg", "pkl/config/example.pkl", "Load Configuration PKL Path File")
	flag.Parse()

	// # Load App Config
	cfg, err := appconfig.LoadFromPath(context.Background(), *cfgPath)
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

	// # Load Application Registry
	dep, repo, serv := app.LoadRegistry(context.Background(), cfg)

	// # Init Go-Chi Router
	router := chi.NewRouter()

	// # Assign Default Global Middleware
	middleware.DefaultGlobal(cfg, dep, router)

	// # Assign Global Middleware
	middleware.Global(cfg, dep, repo, serv, router)

	// # Load Router
	app.LoadRouter(router)

	// # Load HTTP Server API Configuration
	httpServerOption := config.NewHttpServerOption()
	httpServer, err := config.NewHTTPServer(

		// # Required
		//
		// # App Configuration
		//
		cfg,

		// # App Dependency
		//
		dep,

		// # App Repository
		//
		repo,

		// # App Service
		//
		serv,

		// # App Router
		//
		router,

		// # Options
		//
		httpServerOption.WithIdleTimeout(
			time.Duration(cfg.App.Http.IdleTimeout)*time.Second,
		),
		httpServerOption.WithReadHeaderTimeout(
			time.Duration(cfg.App.Http.ReadHeaderTimeout)*time.Second,
		),
		httpServerOption.WithReadTimeout(
			time.Duration(cfg.App.Http.ReadTimeout)*time.Second,
		),
		httpServerOption.WithWriteTimeout(
			time.Duration(cfg.App.Http.WriteTimeout)*time.Second,
		),
	)
	if err != nil {
		dep.Logger.Fatal("Failed to Load Config HTTP Server", "error", err)
	}

	srv, err := httpServer.Create(context.Background())
	if err != nil {
		dep.Logger.Fatal("Failed to Load Initiator HTTP Server", "error", err)
	}

	// # Address
	address := fmt.Sprintf("http://%s:%d", cfg.App.Host, cfg.App.Http.Port)

	go func() {
		dep.Logger.Info(
			"Start HTTP Service API",
			"address", address,
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			dep.Logger.Fatal(
				"Error Start ListenAndServe HTTP Service API",
				"address", address,
				"error", err,
			)
		}
		dep.Logger.Info(
			"Stop HTTP Service API",
			"address", address,
		)
	}()

	<-stopCh

	dep.Logger.Info(
		"Perform shutdown with a maximum timeout of 30 seconds, HTTP Service API",
		"address", address,
	)
	releaseCtx, releaseFn := context.WithTimeout(context.Background(), 30*time.Second)

	defer func() {
		releaseFn()

		// Gracefully Stop Service and Close connection
		if dep.MySqlDB.Loaded() {
			dep.MySqlDB.Value().DB.Close()
		}

		if dep.NatsConn.Loaded() {
			dep.NatsConn.Value().Close()
		}

		dep.Logger.Info(
			"Successfully Stop HTTP Service API, application is exited properly",
			"address", address,
			"error", err,
		)
		if err := dep.LumberjackLogger.Rotate(); err != nil {
			log.Fatalf("rotate file, got error: %+v\n", err)
		}
	}()

	if err := srv.Shutdown(releaseCtx); err != nil {
		dep.Logger.Error(
			"Error Shutdown HTTP Service API, application is exited properly",
			"address", address,
		)
	}
}
