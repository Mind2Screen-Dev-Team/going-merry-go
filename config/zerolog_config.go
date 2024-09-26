package config

import (
	"context"
	"io"
	"os"
	"path"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig/timeformat"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlogger"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zeroLogConfig struct{}

func NewZeroLogConfig() *zeroLogConfig {
	return &zeroLogConfig{}
}

func (z *zeroLogConfig) Loader(ctx context.Context, cfg *appconfig.AppConfig, dep *registry.AppDependency) {
	switch cfg.App.Log.TimeFormat {
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
	if cfg.App.Log.ConsoleLoggingEnabled {
		mw = append(mw, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if cfg.App.Log.FileLoggingEnabled {
		mw = append(mw, dep.LumberjackLogger)
	}

	dep.ZeroLogger = zerolog.New(io.MultiWriter(mw...)).With().Timestamp().Logger()
	dep.Logger = xlogger.NewZeroLogger(&dep.ZeroLogger)
}

func (z *zeroLogConfig) RollingFile(cfg *appconfig.AppConfig) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path.Join(cfg.App.Log.Path, cfg.App.Log.Filename),
		MaxBackups: cfg.App.Log.MaxBackups, // files
		MaxSize:    cfg.App.Log.MaxSize,    // megabytes
		MaxAge:     cfg.App.Log.MaxAge,     // days
		LocalTime:  cfg.App.Log.LocalTime,
	}
}
