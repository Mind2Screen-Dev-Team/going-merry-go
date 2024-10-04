package config

import (
	"context"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	metricNoop "go.opentelemetry.io/otel/metric/noop"
	traceNoop "go.opentelemetry.io/otel/trace/noop"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type OtelParam struct {
	Module        string
	ServerName    string
	ServerAddress string
}

type otelClient struct {
	param OtelParam
}

func NewOtelClient(param OtelParam) *otelClient {
	return &otelClient{param}
}

// Initialize a gRPC connection to be used by both the tracer and meter providers.
func (o *otelClient) NewGrpcConnection(cfg *appconfig.AppConfig) (*grpc.ClientConn, error) {
	if !cfg.Otel.TracerEnabled && !cfg.Otel.MetricEnabled {
		return nil, nil
	}

	// It connects the OpenTelemetry Collector through local gRPC connection.
	// You may replace `localhost:4317` with your endpoint.
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Otel.Grpc.Host, cfg.Otel.Grpc.Port),
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create otel gRPC connection to collector: %w", err)
	}

	return conn, err
}

func (o *otelClient) NewResource(ctx context.Context, cfg *appconfig.AppConfig) (*resource.Resource, error) {
	if !cfg.Otel.TracerEnabled && !cfg.Otel.MetricEnabled {
		return nil, nil
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			semconv.ServiceNameKey.String(o.param.ServerName),
		),
	)
	return res, err
}

// Initializes an OTLP exporter, and configures the corresponding trace provider.
func (o *otelClient) InitTracerProvider(ctx context.Context, cfg *appconfig.AppConfig, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	var (
		tracerShutdownFn func(context.Context) error
		tracerProvider   trace.TracerProvider
	)

	if cfg.Otel.TracerEnabled {
		// Set up a trace exporter
		traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}

		// Register the trace exporter with a TracerProvider, using a batch
		// span processor to aggregate spans before export.
		sp := sdktrace.NewBatchSpanProcessor(traceExporter)
		tracerProvider = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(sp),
		)
		tracerShutdownFn = traceExporter.Shutdown
	} else {
		tracerShutdownFn = func(ctx context.Context) error {
			return nil
		}
		tracerProvider = traceNoop.NewTracerProvider()
	}
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	return tracerShutdownFn, nil
}

// Initializes an OTLP exporter, and configures the corresponding meter provider.
func (o *otelClient) InitMeterProvider(ctx context.Context, cfg *appconfig.AppConfig, res *resource.Resource, conn *grpc.ClientConn) (func(context.Context) error, error) {
	var (
		meterShutdownFn func(context.Context) error
		meterProvider   metric.MeterProvider
	)

	if cfg.Otel.MetricEnabled {
		metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
		}

		meterProvider = sdkmetric.NewMeterProvider(
			sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
			sdkmetric.WithResource(res),
		)
		meterShutdownFn = metricExporter.Shutdown
	} else {
		meterShutdownFn = func(ctx context.Context) error {
			return nil
		}
		meterProvider = metricNoop.NewMeterProvider()
	}
	otel.SetMeterProvider(meterProvider)

	return meterShutdownFn, nil
}

func (o *otelClient) Loader(ctx context.Context, reg *registry.AppRegistry) {
	otelGrpcConn, err := o.NewGrpcConnection(reg.Config)
	if err != nil {
		panic(err)
	}

	otelResource, err := o.NewResource(ctx, reg.Config)
	if err != nil {
		panic(err)
	}

	shutdownTracerProviderFn, err := o.InitTracerProvider(ctx, reg.Config, otelResource, otelGrpcConn)
	if err != nil {
		panic(err)
	}

	shutdownMeterProviderFn, err := o.InitMeterProvider(ctx, reg.Config, otelResource, otelGrpcConn)
	if err != nil {
		panic(err)
	}

	reg.Dependency.OtelModule = o.param.Module
	reg.Dependency.OtelGrpcConn = otelGrpcConn
	reg.Dependency.OtelResource = otelResource
	reg.Dependency.OtelShutdownTracerProviderFn = shutdownTracerProviderFn
	reg.Dependency.OtelShutdownMeterProviderFn = shutdownMeterProviderFn

	reg.Dependency.Tracer = otel.Tracer(o.param.Module,
		trace.WithInstrumentationAttributes(
			attribute.String("server.name", o.param.ServerName),
			attribute.String("server.addr", o.param.ServerAddress),
		),
	)

	reg.Dependency.Metric = otel.Meter(o.param.Module,
		metric.WithInstrumentationAttributes(
			attribute.String("server.name", o.param.ServerName),
			attribute.String("server.addr", o.param.ServerAddress),
		),
	)
}
