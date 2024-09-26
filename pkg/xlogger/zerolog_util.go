package xlogger

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/rs/zerolog"
)

var (
	nopZeroLogger = zerolog.Nop()
	nopLogger     = NewZeroLogger(&nopZeroLogger)
)

func FromReqCtx(ctx context.Context) Logger {
	v, ok := ctx.Value(ctxkey.HTTP_SERVER_APP_LOGGER).(Logger)
	if !ok {
		v = nopLogger
	}
	return v
}
