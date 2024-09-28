package config

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/logconfig/timeformat"

	"github.com/rs/zerolog"
)

type zeroLogConfig struct {
	defaultFields map[string]any
}

func NewZeroLogConfig(defaultFields map[string]any) *zeroLogConfig {
	return &zeroLogConfig{defaultFields}
}

func (z *zeroLogConfig) Loader(ctx context.Context, cfg *appconfig.AppConfig, dep *registry.AppDependency) {
	switch cfg.Log.TimeFormat {
	case timeformat.RFC3339:
		zerolog.TimeFieldFormat = time.RFC3339
	case timeformat.Unix:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	case timeformat.UnixMs:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	case timeformat.UnixMicro:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	case timeformat.UnixNano:
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixNano
	}

	var mw []io.Writer
	if cfg.Log.ConsoleLoggingEnabled {
		mw = append(mw, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.Log.FileLoggingEnabled {
		mw = append(mw, dep.LumberjackLogger)
	}

	dep.ZeroLogger = zerolog.New(io.MultiWriter(mw...)).With().Timestamp().Fields(z.defaultFields).Logger()
}
