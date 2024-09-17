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

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/middleware"
)

func main() {
	// # Load App Config
	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/config/env.pkl")
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
	dep, repo, serv := app.LoadRegistry(cfg)

	// # Init Go-Chi Router
	router := chi.NewRouter()

	// # Assign Default Global Middleware
	middleware.DefaultGlobal(cfg, router)

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

		// # Optional
		httpServerOption.WithIdleTimeout(
			time.Duration(cfg.AppHttp.IdleTimeout)*time.Second,
		),
		httpServerOption.WithReadHeaderTimeout(
			time.Duration(cfg.AppHttp.ReadHeaderTimeout)*time.Second,
		),
		httpServerOption.WithReadTimeout(
			time.Duration(cfg.AppHttp.ReadTimeout)*time.Second,
		),
		httpServerOption.WithWriteTimeout(
			time.Duration(cfg.AppHttp.WriteTimeout)*time.Second,
		),
	)
	if err != nil {
		log.Fatalf("Failed to Load Config HTTP Server, got error %+v\n", err)
	}

	srv, err := httpServer.Create(context.Background())
	if err != nil {
		log.Fatalf("Failed to Load Initiator HTTP Server, got error %+v\n", err)
	}

	go func() {
		log.Printf("Start Service HTTP API on address http://%s:%d\n", cfg.AppHost, cfg.AppHttp.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Service HTTP API on address http://%s:%d, listen and serve : %+v\n", cfg.AppHost, cfg.AppHttp.Port, err)
		}
		log.Printf("Stop Service HTTP API on address http://%s:%d\n", cfg.AppHost, cfg.AppHttp.Port)
	}()

	<-stopCh

	log.Printf("Perform shutdown with a maximum timeout of 30 seconds, Service HTTP API on address http://%s:%d\n", cfg.AppHost, cfg.AppHttp.Port)
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

		log.Printf("Successfully Stop Service HTTP API on address http://%s:%d is exited properly\n", cfg.AppHost, cfg.AppHttp.Port)
	}()

	if err := srv.Shutdown(releaseCtx); err != nil {
		log.Printf("Shutdown Service HTTP API on address http://%s:%d, err : %+v\n", cfg.AppHost, cfg.AppHttp.Port, err)
	}
}
