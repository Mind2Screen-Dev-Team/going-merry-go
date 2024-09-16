package config

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

type HTTPServer struct {
	appConfig     *appconfig.AppConfig
	appDependency *bootstrap.AppDependency

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
	appConfig *appconfig.AppConfig,
	appDependency *bootstrap.AppDependency,
	handler http.Handler,
	opts ...HttpServerOptionFn,
) (*HTTPServer, error) {
	var option httpServerOptionValue

	for _, fn := range opts {
		if err := fn(&option); err != nil {
			return nil, err
		}
	}

	return &HTTPServer{appConfig, appDependency, handler, &option}, nil
}

func (h *HTTPServer) Create(_ context.Context) (*http.Server, error) {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", h.appConfig.AppHost, h.appConfig.AppHttp.Port),
		Handler:           h.handler,
		IdleTimeout:       h.option.IdleTimeout,
		ReadHeaderTimeout: h.option.ReadHeaderTimeout,
		ReadTimeout:       h.option.ReadTimeout,
		WriteTimeout:      h.option.WriteTimeout,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(context.Background(), ctxkey.CTX_KEY_HTTP_SERVER_APP_CONFIG, h.appConfig)
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
