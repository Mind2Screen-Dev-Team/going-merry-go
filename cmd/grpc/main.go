package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	icpt_stream "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/stream"
	icpt_unary "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/unary"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	APP_NAME = "grpc-api"
)

func main() {
	// # Parse App Config Path
	cfgPath := flag.String("cfg", "./example.pkl", "Load Configuration PKL Path File")
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

	address := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)

	// # Load Application Registry
	reg := app.LoadRegistry(context.Background(), cfg, app.DependencyRegistryLoaderParams{
		Module:      "grpc.api.app",
		LogFilename: fmt.Sprintf("%s.log", cfg.Grpc.ServiceName),
		LogDefaultFields: map[string]any{
			"serviceName":    cfg.Grpc.ServiceName,
			"serviceAddress": address,
		},
	})

	// # Must Load Dependency At Startup
	if err := app.Startup(APP_NAME, reg); err != nil {
		panic(err)
	}

	reg.InterruptContext = interruptCtx

	// # Init Logger
	logger := xlogger.NewZeroLogger(&reg.Dependency.ZeroLogger)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("Start Listen GRPC Service API", "error", err)
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
		grpc.ChainStreamInterceptor(
			icpt_stream.RegisterRegistry(reg),
			icpt_stream.RequestIDInterceptor(),
			icpt_stream.Logging(reg.Dependency.ZeroLogger),
		),
		grpc.ChainUnaryInterceptor(
			icpt_unary.RegisterRegistry(reg),
			icpt_unary.RequestIDInterceptor(),
			icpt_unary.Logging(reg.Dependency.ZeroLogger),
		),
	)

	// Grpc Server
	server := grpc.NewServer(opts...)

	// Grpc Service Loader
	app.AppGrpcServiceLoader(server, reg)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	go func() {
		logger.Info("Start GRPC Service API")
		if err := server.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			logger.Fatal("Error Start Serve GRPC Service API", "error", err)
		}
		logger.Info("Stop GRPC Service API")
	}()

	<-interruptCtx.Done()

	// Gracefully Stop Service and Close connection
	defer app.Shutdown(APP_NAME, reg)

	// # Wait to Shutdown GRPC Server
	logger.Info("Perform graceful stop, GRPC Service API")
	server.GracefulStop()

}
