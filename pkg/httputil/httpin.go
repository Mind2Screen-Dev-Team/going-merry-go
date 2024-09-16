package httputil

import (
	"context"
	"net/http"

	"github.com/ggicci/httpin"
)

func ParseInput[T any](ctx context.Context) *T {
	input, ok := ctx.Value(httpin.Input).(*T)
	if !ok {
		return nil
	}
	return input
}

// Middleware to Parse Request input
func WithInput[T any]() func(http.Handler) http.Handler {
	var in T
	return httpin.NewInput(&in)
}
