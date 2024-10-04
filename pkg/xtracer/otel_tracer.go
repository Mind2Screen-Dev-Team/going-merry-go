package xtracer

import (
	"context"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/config"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewTracerHandlerFunc(handlerFn http.HandlerFunc, operation string, opts ...otelhttp.Option) http.HandlerFunc {
	return otelhttp.NewHandler(
		http.HandlerFunc(handlerFn),
		operation,
		opts...,
	).ServeHTTP
}

func NewTracerHandler(handlerFn http.HandlerFunc, operation string, opts ...otelhttp.Option) http.Handler {
	return otelhttp.NewHandler(
		http.HandlerFunc(handlerFn),
		operation,
		opts...,
	)
}

func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	var requestId string
	if v, ok := ctx.Value(ctxkey.RequestIDKey).(string); ok {
		requestId = v
	}

	opts = append(opts, trace.WithTimestamp(time.Now()), trace.WithAttributes(
		attribute.String("requestId", requestId),
	))

	return config.LoadAppTracer(ctx).Start(ctx, spanName, opts...)
}
