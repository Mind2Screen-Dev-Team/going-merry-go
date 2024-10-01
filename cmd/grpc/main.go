package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
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

	address := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

	// # Load Application Registry
	dep, repo, serv, prov := app.LoadRegistry(context.Background(), cfg, app.AppDependencyLoaderParams{
		LogFilename: fmt.Sprintf("%s.log", cfg.Grpc.ServiceName),
		LogDefaultFields: map[string]any{
			"serviceName":    cfg.Grpc.ServiceName,
			"serviceAddress": address,
		},
	})

	// # Must Load Dependency At Startup
	if err := app.MustLoadDependencyAtStartup("grpc-api", cfg, dep); err != nil {
		panic(err)
	}

	// # Init Logger
	logger := xlogger.NewZeroLogger(&dep.ZeroLogger)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Info("Start Listen GRPC Service API", "error", err)
	}

	var opts []grpc.ServerOption

	if cfg.Grpc.KeepAlive.Enabled {
		opts = append(opts,
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
				MinTime:             time.Duration(cfg.Grpc.KeepAlive.Policy.MinTime) * time.Second,
				PermitWithoutStream: cfg.Grpc.KeepAlive.Policy.PermitWithoutStream,
			}),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle:     time.Duration(cfg.Grpc.KeepAlive.Parameter.MaxConnectionIdle) * time.Second,
				MaxConnectionAge:      time.Duration(cfg.Grpc.KeepAlive.Parameter.MaxConnectionAge) * time.Second,
				MaxConnectionAgeGrace: time.Duration(cfg.Grpc.KeepAlive.Parameter.MaxConnectionAgeGrace) * time.Second,
				Time:                  time.Duration(cfg.Grpc.KeepAlive.Parameter.Time) * time.Second,
				Timeout:               time.Duration(cfg.Grpc.KeepAlive.Parameter.Timeout) * time.Second,
			}),
		)
	}

	// # Set GRPC Unary / Stream Interceptors
	opts = append(opts,
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)

	server := grpc.NewServer(opts...)

	// GRPC Handler Loader
	app.AppGRPCHandlerLoader(cfg, server, dep, repo, prov, serv)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	go func() {
		logger.Info("Start GRPC Service API")
		if err := server.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			logger.Fatal("Error Start Serve GRPC Service API", "error", err)
		}
		logger.Info("Stop GRPC Service API")
	}()

	<-stopCh

	// Gracefully Stop Service and Close connection
	defer func() {
		if dep.MySqlDB.Loaded() {
			dep.MySqlDB.Value().DB.Close()
			logger.Info("Successfully Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port), "mysqlDB", cfg.Mysql.Db)
		}

		if dep.NatsConn.Loaded() {
			dep.NatsConn.Value().Close()
			logger.Info("Successfully Close Nats Connection", "natsAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port))
		}

		logger.Info("Successfully gracefully stop GRPC Service API, application is exited properly")
		if err := dep.LumberjackLogger.Rotate(); err != nil {
			log.Fatalf("grpc-api rotate logging file, got error: %+v\n", err)
		}
	}()

	// # Wait to Shutdown GRPC Server
	logger.Info("Perform graceful stop, GRPC Service API")
	server.GracefulStop()

}
