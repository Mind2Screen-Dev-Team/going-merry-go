package config

import (
	"context"
	"path"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"gopkg.in/natefinch/lumberjack.v2"
)

type lumberJackConfig struct{}

func NewLumberJackConfig() *lumberJackConfig {
	return &lumberJackConfig{}
}

func (lc *lumberJackConfig) Loader(ctx context.Context, cfg *appconfig.AppConfig, dep *registry.AppDependency) {
	dep.LumberjackLogger = &lumberjack.Logger{
		Filename:   path.Join(cfg.App.Log.Path, cfg.App.Log.Filename),
		MaxBackups: cfg.App.Log.MaxBackups, // how much files
		MaxSize:    cfg.App.Log.MaxSize,    // how much megabytes
		MaxAge:     cfg.App.Log.MaxAge,     // how much days
		LocalTime:  cfg.App.Log.LocalTime,  // default UTC
		Compress:   cfg.App.Log.Compress,   // default false
	}
}
