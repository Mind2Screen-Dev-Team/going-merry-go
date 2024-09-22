package config

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

type HTTPServer struct {
	cfg  *appconfig.AppConfig
	dep  *registry.AppDependency
	repo *registry.AppRepository
	serv *registry.AppService

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
	cfg *appconfig.AppConfig,
	dep *registry.AppDependency,
	repo *registry.AppRepository,
	serv *registry.AppService,
	handler http.Handler,
	opts ...HttpServerOptionFn,
) (*HTTPServer, error) {
	var option httpServerOptionValue

	for _, fn := range opts {
		if err := fn(&option); err != nil {
			return nil, err
		}
	}

	return &HTTPServer{
		cfg,
		dep,
		repo,
		serv,
		handler,
		&option,
	}, nil
}

func (h *HTTPServer) Create(_ context.Context) (*http.Server, error) {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", h.cfg.App.Host, h.cfg.App.Http.Port),
		Handler:           h.handler,
		IdleTimeout:       h.option.IdleTimeout,
		ReadHeaderTimeout: h.option.ReadHeaderTimeout,
		ReadTimeout:       h.option.ReadTimeout,
		WriteTimeout:      h.option.WriteTimeout,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.Background()

			// # Assign a Value To Context
			ctx = context.WithValue(ctx, ctxkey.HTTP_SERVER_APP_CONFIG, h.cfg)
			ctx = context.WithValue(ctx, ctxkey.HTTP_SERVER_APP_DEPENDENCY, h.dep)
			ctx = context.WithValue(ctx, ctxkey.HTTP_SERVER_APP_REPOSITORY, h.repo)
			ctx = context.WithValue(ctx, ctxkey.HTTP_SERVER_APP_SERVICE, h.serv)
			ctx = context.WithValue(ctx, ctxkey.HTTP_SERVER_APP_LOGGER, h.dep.Logger)

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
