package app

import (
	"fmt"
	"log"
	"slices"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"
)

func Shutdown(service string, reg *registry.AppRegistry, fns ...func()) {
	// # Init Logger
	logger := xlogger.NewZeroLogger(&reg.Dependency.ZeroLogger)

	for _, fn := range fns {
		fn()
	}

	if slices.Contains([]string{"rest-api", "grpc-api"}, service) {
		if err := reg.Dependency.OtelShutdownTracerProviderFn(reg.InterruptContext); err != nil {
			logger.Error("failed to shutdown tracer provider", "error", err)
		}

		if err := reg.Dependency.OtelShutdownMeterProviderFn(reg.InterruptContext); err != nil {
			logger.Error("failed to shutdown meter provider", "error", err)
		}

		if reg.Dependency.OtelGrpcConn != nil {
			if err := reg.Dependency.OtelGrpcConn.Close(); err != nil {
				logger.Error("Error Close Otel GRPC Connection", "otelGrpcAddr", fmt.Sprintf("%s:%d", reg.Config.Otel.Grpc.Host, reg.Config.Otel.Grpc.Port), "error", err)
			} else {
				logger.Info("Successfully Close Otel GRPC Connection", "otelGrpcAddr", fmt.Sprintf("%s:%d", reg.Config.Otel.Grpc.Host, reg.Config.Otel.Grpc.Port))
			}
		}

		if reg.Dependency.MySqlDB.Loaded() {
			if err := reg.Dependency.MySqlDB.Value().DB.Close(); err != nil {
				logger.Error("Error Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", reg.Config.Mysql.Host, reg.Config.Mysql.Port), "mysqlDB", reg.Config.Mysql.Db, "error", err)
			} else {
				logger.Info("Successfully Close MySQL DB Connection", "mysqlAddr", fmt.Sprintf("%s:%d", reg.Config.Mysql.Host, reg.Config.Mysql.Port), "mysqlDB", reg.Config.Mysql.Db)
			}
		}

		if reg.Dependency.NatsConn.Loaded() {
			reg.Dependency.NatsConn.Value().Close()
			logger.Info("Successfully Close Nats Connection", "natsAddr", fmt.Sprintf("%s:%d", reg.Config.Mysql.Host, reg.Config.Mysql.Port))
		}

		logger.Info("Successfully gracefuly Stop HTTP Service API, application is exited properly")
		if err := reg.Dependency.LumberjackLogger.Rotate(); err != nil {
			log.Fatalf("%q service rotate logging file, got error: %+v\n", service, err)
		}
	}
}
