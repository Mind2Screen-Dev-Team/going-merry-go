package config

import (
	"context"
	"net/http"
	"time"
)

func NewHttpServer(addr string, mux http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
	}
}

type HTTPServer struct {
	address string
	mux     *http.ServeMux
	option  *httpServerOptionValue
}

type httpServerOptionValue struct {
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
}

func NewHTTPServer(address string, mux *http.ServeMux, opts ...HttpServerOptionFn) (*HTTPServer, error) {
	var option httpServerOptionValue

	for _, opFn := range opts {
		if err := opFn(&option); err != nil {
			return nil, err
		}
	}

	return &HTTPServer{address, mux, &option}, nil
}

func (h *HTTPServer) Create(_ context.Context) (*http.Server, error) {
	return &http.Server{
		Addr:              h.address,
		Handler:           h.mux,
		IdleTimeout:       h.option.IdleTimeout,
		ReadHeaderTimeout: h.option.ReadHeaderTimeout,
		ReadTimeout:       h.option.ReadTimeout,
		WriteTimeout:      h.option.WriteTimeout,
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
