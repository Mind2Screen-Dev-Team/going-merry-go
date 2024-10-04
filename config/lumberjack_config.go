package config

import (
	"context"
	"path"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"

	"gopkg.in/natefinch/lumberjack.v2"
)

type lumberJackConfig struct {
	filename string
}

func NewLumberJackConfig(filename string) *lumberJackConfig {
	return &lumberJackConfig{filename}
}

func (lc *lumberJackConfig) Loader(ctx context.Context, reg *registry.AppRegistry) error {
	reg.Dependency.LumberjackLogger = &lumberjack.Logger{
		Filename:   path.Join(reg.Config.Log.Path, lc.filename),
		MaxBackups: reg.Config.Log.MaxBackups, // how much backup files
		MaxSize:    reg.Config.Log.MaxSize,    // how much maximum megabytes
		MaxAge:     reg.Config.Log.MaxAge,     // how much maximum days, default is 0 that means not deleted old logs
		LocalTime:  reg.Config.Log.LocalTime,  // default UTC
		Compress:   reg.Config.Log.Compress,   // default false
	}
	return nil
}
