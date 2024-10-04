package icpt_util

import (
	"context"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/constant/ctxkey"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/rs/zerolog"
)

func Logger(l zerolog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		// Trace ID
		fields = append(fields, "requestId", ctx.Value(ctxkey.RequestIDKey))

		switch lvl {
		case logging.LevelDebug:
			l.Debug().Fields(fields).Msg(msg)
		case logging.LevelInfo:
			l.Info().Fields(fields).Msg(msg)
		case logging.LevelWarn:
			l.Warn().Fields(fields).Msg(msg)
		case logging.LevelError:
			l.Error().Fields(fields).Msg(msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
