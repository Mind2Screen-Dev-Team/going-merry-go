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
	interceptor_unary "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/grpc/interceptor/unary"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

// InterceptorLogger adapts zerolog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		md, _ := metadata.FromIncomingContext(ctx)

		// Trace ID
		traceId := md.Get("traceId")
		if len(traceId) > 0 {
			fields = append(fields, "traceId", traceId[0])
		}

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Fields(fields).Msg(msg)
		case logging.LevelInfo:
			l.Info().Fields(fields).Msg(msg)
		case logging.LevelWarn:
			l.Warn().Fields(fields).Msg(msg)
		case logging.LevelError:
			l.Error().Fields(fields).Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

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
	reg := app.LoadRegistry(context.Background(), cfg, app.AppDependencyLoaderParams{
		LogFilename: fmt.Sprintf("%s.log", cfg.Grpc.ServiceName),
		LogDefaultFields: map[string]any{
			"serviceName":    cfg.Grpc.ServiceName,
			"serviceAddress": address,
		},
	})

	// # Must Load Dependency At Startup
	if err := app.MustLoadDependencyAtStartup("grpc-api", reg); err != nil {
		panic(err)
	}

	// # Init Logger
	logger := xlogger.NewZeroLogger(&reg.Dependency.ZeroLogger)

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

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	// # Set GRPC Unary / Stream Interceptors
	opts = append(opts,
		grpc.ChainUnaryInterceptor(
			interceptor_unary.RequestIDInterceptor(),
			logging.UnaryServerInterceptor(InterceptorLogger(reg.Dependency.ZeroLogger), loggingOpts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(reg.Dependency.ZeroLogger), loggingOpts...),
		),
	)

	server := grpc.NewServer(opts...)

	// GRPC Handler Loader
	app.AppGRPCHandlerLoader(server, reg)

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
		if reg.Dependency.MySqlDB.Loaded() {
			reg.Dependency.MySqlDB.Value().DB.Close()
			logger.Info("Successfully Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port), "mysqlDB", cfg.Mysql.Db)
		}

		if reg.Dependency.NatsConn.Loaded() {
			reg.Dependency.NatsConn.Value().Close()
			logger.Info("Successfully Close Nats Connection", "natsAddr", fmt.Sprintf("%s:%d", cfg.Mysql.Host, cfg.Mysql.Port))
		}

		logger.Info("Successfully gracefully stop GRPC Service API, application is exited properly")
		if err := reg.Dependency.LumberjackLogger.Rotate(); err != nil {
			log.Fatalf("grpc-api rotate logging file, got error: %+v\n", err)
		}
	}()

	// # Wait to Shutdown GRPC Server
	logger.Info("Perform graceful stop, GRPC Service API")
	server.GracefulStop()

}
