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

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/http/middleware"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	"github.com/go-chi/chi/v5"
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
	interruptCtx, interruptFn := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	defer interruptFn()

	addr := fmt.Sprintf("http://%s:%d", cfg.Http.Host, cfg.Http.Port)

	// # Load Application Registry
	reg := app.LoadRegistry(interruptCtx, cfg, app.DependencyRegistryLoaderParams{
		Module:      "rest.api.app",
		ServerName:  cfg.Http.ServiceName,
		ServerAddr:  addr,
		LogFilename: fmt.Sprintf("%s.log", cfg.Http.ServiceName),
		LogDefaultFields: map[string]any{
			"serviceName":    cfg.Http.ServiceName,
			"serviceAddress": addr,
			"servicePID":     os.Getpid(),
		},
	})

	// # Must Load Dependency At Startup
	if err := app.MustLoadDependencyAtStartup("rest-api", reg); err != nil {
		panic(err)
	}

	// # Init Logger
	logger := xlogger.NewZeroLogger(&reg.Dependency.ZeroLogger)

	// # Init Go-Chi Router
	router := chi.NewRouter()

	// # Assign Default Global Middleware
	middleware.DefaultGlobal(reg, router)

	// # Assign Global Middleware
	middleware.Global(reg, router)

	// # Load Router
	app.LoadRouter(router)

	// # Load HTTP Server API Configuration
	httpServerOption := config.NewHttpServerOption()
	httpServer, err := config.NewHTTPServer(

		// # App Registry
		//
		reg,

		// # App Router
		//
		router,

		// # Options
		//
		httpServerOption.WithIdleTimeout(
			time.Duration(cfg.Http.IdleTimeout)*time.Second,
		),
		httpServerOption.WithReadHeaderTimeout(
			time.Duration(cfg.Http.ReadHeaderTimeout)*time.Second,
		),
		httpServerOption.WithReadTimeout(
			time.Duration(cfg.Http.ReadTimeout)*time.Second,
		),
		httpServerOption.WithWriteTimeout(
			time.Duration(cfg.Http.WriteTimeout)*time.Second,
		),
	)
	if err != nil {
		logger.Fatal("Failed to Load Config HTTP Server", "error", err)
	}

	srv, err := httpServer.Create(interruptCtx)
	if err != nil {
		logger.Fatal("Failed to Load Initiator HTTP Server", "error", err)
	}

	go func() {
		logger.Info("Start HTTP Service API")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error Start ListenAndServe HTTP Service API", "error", err)
		}
		logger.Info("Stop HTTP Service API")
	}()

	<-interruptCtx.Done()

	logger.Info("Perform shutdown with a maximum timeout of 30 seconds, HTTP Service API")
	releaseCtx, releaseFn := context.WithTimeout(context.Background(), 30*time.Second)

	// Gracefully Stop Service and Close connection
	defer func() {
		releaseFn()

		if err := reg.Dependency.OtelShutdownTracerProviderFn(interruptCtx); err != nil {
			logger.Error("failed to shutdown tracer provider", "error", err)
		}

		if err := reg.Dependency.OtelShutdownMeterProviderFn(interruptCtx); err != nil {
			logger.Error("failed to shutdown meter provider", "error", err)
		}

		if reg.Dependency.OtelGrpcConn != nil {
			if err := reg.Dependency.OtelGrpcConn.Close(); err != nil {
				logger.Error("Error Close Otel GRPC Connection", "otelGrpcAddr", fmt.Sprintf("%s:%d", cfg.Otel.Grpc.Host, cfg.Otel.Grpc.Port), "error", err)
			} else {
				logger.Info("Successfully Close Otel GRPC Connection", "otelGrpcAddr", fmt.Sprintf("%s:%d", cfg.Otel.Grpc.Host, cfg.Otel.Grpc.Port))
			}
		}

		if reg.Dependency.MySqlDB.Loaded() {
			if err := reg.Dependency.MySqlDB.Value().DB.Close(); err != nil {
				logger.Error("Error Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port), "mysqlDB", cfg.Mysql.Db, "error", err)
			} else {
				logger.Info("Successfully Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port), "mysqlDB", cfg.Mysql.Db)
			}
		}

		if reg.Dependency.NatsConn.Loaded() {
			reg.Dependency.NatsConn.Value().Close()
			logger.Info("Successfully Close Nats Connection", "natsAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port))
		}

		logger.Info("Successfully gracefuly Stop HTTP Service API, application is exited properly")
		if err := reg.Dependency.LumberjackLogger.Rotate(); err != nil {
			log.Fatalf("rest-api service rotate logging file, got error: %+v\n", err)
		}
	}()

	if err := srv.Shutdown(releaseCtx); err != nil {
		logger.Error("Error Shutdown HTTP Service API, application is exited properly", "error", err)
	}
}
