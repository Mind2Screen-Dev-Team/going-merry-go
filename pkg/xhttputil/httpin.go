package xhttputil

import (
	"context"
	"net/http"

	"github.com/ggicci/httpin"
)

// Utility to load input from request context. for example:
//
//	func SomeHandler(w http.ResponseWriter, r *http.Request) {
//		ctx	:= r.Context()
//		// Output: *dto.AuthLoginReqDTO
//		input := httputil.LoadInput[dto.AuthLoginReqDTO](ctx)
//	}
//
func LoadInput[T any](ctx context.Context) *T {
	input, ok := ctx.Value(httpin.Input).(*T)
	if !ok {
		var in T
		return &in
	}
	return input
}

// Middleware to parse request input: path variable, query params, body (json,xml,...) and form. for example:
// 	
//	router.Use(httputil.WithInput[dto.AuthLoginReqDTO]())
//
func WithInput[T any]() func(http.Handler) http.Handler {
	var in T
	return httpin.NewInput(&in)
}
