package xhttputil

import (
	"context"
	"net/http"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xtracer"
	"github.com/ggicci/httpin"
	"github.com/ggicci/httpin/core"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/guregu/null.v4"
)

// Utility to load input from request context. for example:
//
//	func SomeHandler(w http.ResponseWriter, r *http.Request) {
//		ctx	:= r.Context()
//		// Output: *dto.AuthLoginReqDTO
//		input := httputil.LoadInput[dto.AuthLoginReqDTO](ctx)
//	}
func LoadInput[T any](ctx context.Context) *T {
	var span trace.Span

	ctx, span = xtracer.Start(ctx, "load.input.from.request")
	defer span.End()

	input, ok := ctx.Value(httpin.Input).(*T)
	if !ok {
		var in T
		return &in
	}

	return input
}

type InputOption interface {
	// # Ignored because this already set on default config.
	// WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request, err error)) InputOptionFn

	// Otel span operation name
	WithOperationName(n string) InputOptionFn

	/*
		Understanding Max Memory Allocation, i.e:
			- 1 kilobyte (KB) = 1024 bytes;
			- 1 megabyte (MB) = 1024 kilobytes;
			- 1 MB = 1024 * 1024 = 1,048,576 bytes;
			- 1 MB = 1 * 1024 * 1024;
			- formula for: X MB = X * 1024 * 1024;
	*/
	WithMaxMemory(n int64) InputOptionFn
	WithNestedDirectivesEnabled(enabled bool) InputOptionFn
}

type InputOptionFn func(i *inputOptionValues)

type inputOptionValues struct {
	// # Ignored because this already set on default config.
	// errorHandler func(w http.ResponseWriter, r *http.Request, err error)

	operation               null.String
	maxMemory               null.Int
	nestedDirectivesEnabled null.Bool
}

type inputOptions struct{}

func NewInputOption() InputOption {
	return inputOptions{}
}

// # Ignored because this already set on default config.
// func (inputOptions) WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request, err error)) InputOptionFn {
// 	return func(i *inputOptionValues) {
// 		i.errorHandler = fn
// 	}
// }

func (inputOptions) WithOperationName(n string) InputOptionFn {
	return func(i *inputOptionValues) {
		i.operation = null.NewString(n, true)
	}
}

func (inputOptions) WithMaxMemory(n int64) InputOptionFn {
	return func(i *inputOptionValues) {
		i.maxMemory = null.NewInt(n, true)
	}
}

func (inputOptions) WithNestedDirectivesEnabled(enabled bool) InputOptionFn {
	return func(i *inputOptionValues) {
		i.nestedDirectivesEnabled = null.NewBool(enabled, true)
	}
}

// Middleware to parse request input: path variables, headers, query params, body (json and xml) and forms. for example:
//
//	router.Use(httputil.WithInput[dto.AuthLoginReqDTO]())
func WithInput[T any](opts ...InputOptionFn) func(http.Handler) http.Handler {
	var opt inputOptionValues
	for _, optFn := range opts {
		optFn(&opt)
	}

	var copts []core.Option

	// # Ignored because this already set on default config.
	// if opt.errorHandler != nil {
	// 	copts = append(copts, core.WithErrorHandler(opt.errorHandler))
	// }

	if opt.maxMemory.Valid {
		copts = append(copts, core.WithMaxMemory(opt.maxMemory.Int64))
	}
	if opt.nestedDirectivesEnabled.Valid {
		copts = append(copts, core.WithNestedDirectivesEnabled(opt.nestedDirectivesEnabled.Bool))
	}

	var in T
	return func(next http.Handler) http.Handler {
		return otelhttp.NewMiddleware(opt.operation.String)(
			httpin.NewInput(&in, copts...)(next),
		)
	}
}
