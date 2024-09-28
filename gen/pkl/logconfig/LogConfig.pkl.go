// Code generated from Pkl module `LogConfig`. DO NOT EDIT.
package logconfig

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/logconfig/timeformat"
	"github.com/apple/pkl-go/pkl"
)

type LogConfig struct {
	// Log Path
	Path string `pkl:"path"`

	// Log Max Backups
	MaxBackups int `pkl:"maxBackups"`

	// Log Max Size, in Mega Bytes (X MB)
	MaxSize int `pkl:"maxSize"`

	// Log Max Age for backup will deleted, this value in days when 0 not deleted old backup
	MaxAge int `pkl:"maxAge"`

	// Log Use to Local Time, default UTC
	LocalTime bool `pkl:"localTime"`

	// Log Timestamp Used
	TimeFormat timeformat.TimeFormat `pkl:"timeFormat"`

	// Log Write to Console
	ConsoleLoggingEnabled bool `pkl:"consoleLoggingEnabled"`

	// Log Write to File
	FileLoggingEnabled bool `pkl:"fileLoggingEnabled"`

	// Log Rotation Compress
	Compress bool `pkl:"compress"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a LogConfig
func LoadFromPath(ctx context.Context, path string) (ret *LogConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a LogConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*LogConfig, error) {
	var ret LogConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
