package config

import (
	"context"
	"path"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"

	"gopkg.in/natefinch/lumberjack.v2"
)

type lumberJackConfig struct {
	filename string
}

func NewLumberJackConfig(filename string) *lumberJackConfig {
	return &lumberJackConfig{filename}
}

func (lc *lumberJackConfig) Loader(ctx context.Context, cfg *appconfig.AppConfig, dep *registry.AppDependency) {
	dep.LumberjackLogger = &lumberjack.Logger{
		Filename:   path.Join(cfg.Log.Path, lc.filename),
		MaxBackups: cfg.Log.MaxBackups, // how much backup files
		MaxSize:    cfg.Log.MaxSize,    // how much maximum megabytes
		MaxAge:     cfg.Log.MaxAge,     // how much maximum days, default is 0 that means not deleted old logs
		LocalTime:  cfg.Log.LocalTime,  // default UTC
		Compress:   cfg.Log.Compress,   // default false
	}
}
