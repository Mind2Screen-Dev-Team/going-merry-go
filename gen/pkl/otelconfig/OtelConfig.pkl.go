// Code generated from Pkl module `OtelConfig`. DO NOT EDIT.
package otelconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type OtelConfig struct {
	// Otel Tracer is enabled
	TracerEnabled bool `pkl:"tracerEnabled"`

	// Otel Metric is enabled
	MetricEnabled bool `pkl:"metricEnabled"`

	// Otel GRPC
	Grpc *Grpc `pkl:"grpc"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a OtelConfig
func LoadFromPath(ctx context.Context, path string) (ret *OtelConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a OtelConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*OtelConfig, error) {
	var ret OtelConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
