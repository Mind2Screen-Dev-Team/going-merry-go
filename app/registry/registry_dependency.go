package registry

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// # DEPENDENCY

type DependencyRegistry struct {
	// register your dependency on here
	ZeroLogDefaultFields map[string]any
	LumberjackLogger     *lumberjack.Logger
	ZeroLogger           zerolog.Logger

	OtelModule                   string
	OtelGrpcConn                 *grpc.ClientConn
	OtelResource                 *resource.Resource
	OtelShutdownTracerProviderFn func(context.Context) error
	OtelShutdownMeterProviderFn  func(context.Context) error
	Tracer                       trace.Tracer
	Metric                       metric.Meter

	Storage           xlazy.Loader[*minio.Client]
	MySqlDB           xlazy.Loader[*sqlx.DB]
	Redis             xlazy.Loader[*redis.Client]
	NatsConn          xlazy.Loader[*nats.Conn]
	NatsJetStreamConn xlazy.Loader[jetstream.JetStream]
}

func NewDependencyRegistry() *DependencyRegistry {
	return &DependencyRegistry{}
}
