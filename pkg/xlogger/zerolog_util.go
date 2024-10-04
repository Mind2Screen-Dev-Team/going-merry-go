package xlogger

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/rs/zerolog"
)

var (
	nopZeroLogger = zerolog.Nop()
)

func FromReqCtx(ctx context.Context) Logger {
	v, ok := ctx.Value(ctxkey.REGISTRY_APP_LOGGER).(*zerolog.Logger)
	if !ok {
		v = &nopZeroLogger
	}
	return NewZeroLogger(v)
}
