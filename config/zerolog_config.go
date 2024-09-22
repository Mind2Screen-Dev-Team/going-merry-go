package config

import (
	"context"
	"io"
	"os"
	"path"
	"time"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/rs/zerolog"

	"gopkg.in/natefinch/lumberjack.v2"
)

type zeroLogConfig struct{}

func NewZeroLogConfig() *zeroLogConfig {
	return &zeroLogConfig{}
}

func (z *zeroLogConfig) Loader(ctx context.Context, cfg *appconfig.AppConfig, app *registry.AppDependency) {
	multi := io.MultiWriter(
		zerolog.ConsoleWriter{
			NoColor:    true,
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		},
		z.RollingFile(cfg),
	)
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()
	app.Logger = &logger
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
