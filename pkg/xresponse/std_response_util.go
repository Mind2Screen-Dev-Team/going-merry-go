package xresponse

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
)

func copyCtxHttpServerValue(dst context.Context, src context.Context, key ctxkey.CtxKeyHttpServer) context.Context {
	return context.WithValue(dst, key, src.Value(key))
}
