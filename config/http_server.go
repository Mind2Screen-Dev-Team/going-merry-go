package config

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
)

type HTTPServer struct {
	reg     *registry.AppRegistry
	handler http.Handler
	option  *httpServerOptionValue
}

type httpServerOptionValue struct {
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
}

func NewHTTPServer(
	reg *registry.AppRegistry,
	handler http.Handler,
	opts ...HttpServerOptionFn,
) (*HTTPServer, error) {
	var option httpServerOptionValue

	for _, fn := range opts {
		if err := fn(&option); err != nil {
			return nil, err
		}
	}

	return &HTTPServer{reg, handler, &option}, nil
}

func (h *HTTPServer) Create(ctx context.Context) (*http.Server, error) {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", h.reg.Config.Http.Host, h.reg.Config.Http.Port),
		Handler:           h.handler,
		IdleTimeout:       h.option.IdleTimeout,
		ReadHeaderTimeout: h.option.ReadHeaderTimeout,
		ReadTimeout:       h.option.ReadTimeout,
		WriteTimeout:      h.option.WriteTimeout,
		BaseContext: func(l net.Listener) context.Context {

			// # Assign a Value To Context
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_CONFIG, h.reg.Config)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_DEPENDENCY, h.reg.Dependency)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_REPOSITORY, h.reg.Repository)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_PROVIDER, h.reg.Provider)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_SERVICE, h.reg.Service)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_LOGGER, &h.reg.Dependency.ZeroLogger)
			ctx = context.WithValue(ctx, ctxkey.REGISTRY_APP_TRACER, h.reg.Dependency.Tracer)

			return ctx
		},
	}, nil
}

// # HTTP Server OPTIONS

type HttpServerOptionFn func(in *httpServerOptionValue) error

type HttpServerOption struct{}

func NewHttpServerOption() HttpServerOption {
	return HttpServerOption{}
}

func (HttpServerOption) WithIdleTimeout(value time.Duration) HttpServerOptionFn {
	return func(in *httpServerOptionValue) error {
		in.IdleTimeout = value
		return nil
	}
}

func (HttpServerOption) WithReadHeaderTimeout(value time.Duration) HttpServerOptionFn {
	return func(in *httpServerOptionValue) error {
		in.ReadHeaderTimeout = value
		return nil
	}
}

func (HttpServerOption) WithReadTimeout(value time.Duration) HttpServerOptionFn {
	return func(in *httpServerOptionValue) error {
		in.ReadTimeout = value
		return nil
	}
}

func (HttpServerOption) WithWriteTimeout(value time.Duration) HttpServerOptionFn {
	return func(in *httpServerOptionValue) error {
		in.WriteTimeout = value
		return nil
	}
}
