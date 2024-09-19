package xhttputil

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/ggicci/httpin/core"
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
	input, ok := ctx.Value(httpin.Input).(*T)
	if !ok {
		var in T
		return &in
	}
	return input
}

type InputOption interface {
	WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request, err error)) InputOptionFn
	WithMaxMemory(n int64) InputOptionFn
	WithNestedDirectivesEnabled(enabled bool) InputOptionFn
}

type InputOptionFn func(i *inputOptionValues)

type inputOptionValues struct {
	errorHandler            func(w http.ResponseWriter, r *http.Request, err error)
	maxMemory               null.Int
	nestedDirectivesEnabled null.Bool
}

type inputOptions struct{}

func NewInputOption() InputOption {
	return inputOptions{}
}

func (inputOptions) WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request, err error)) InputOptionFn {
	return func(i *inputOptionValues) {
		i.errorHandler = fn
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
	if opt.errorHandler != nil {
		copts = append(copts, core.WithErrorHandler(opt.errorHandler))
	}
	if opt.maxMemory.Valid {
		copts = append(copts, core.WithMaxMemory(opt.maxMemory.Int64))
	}
	if opt.nestedDirectivesEnabled.Valid {
		copts = append(copts, core.WithNestedDirectivesEnabled(opt.nestedDirectivesEnabled.Bool))
	}

	var in T
	return httpin.NewInput(&in, copts...)
}

func DeepCopyRequest(r *http.Request) *http.Request {
	// Read the body if it's non-nil
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
		// Refill the original request body to preserve it for further usage.
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Create a shallow copy of the request
	rCopy := r.Clone(r.Context())

	// Replace the body of the new request with a new reader wrapping the copied bytes
	if bodyBytes != nil {
		rCopy.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return rCopy
}
