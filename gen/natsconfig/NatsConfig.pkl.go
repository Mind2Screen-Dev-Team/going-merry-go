// Code generated from Pkl module `NatsConfig`. DO NOT EDIT.
package natsconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type NatsConfig struct {
	// Whether Nats connections are enabled or not.
	Enabled bool `pkl:"enabled"`

	// The hostname that Nats listens on
	Host string `pkl:"host"`

	// The port that Nats listens on
	Port int `pkl:"port"`

	// Authorization settings for Nats
	Auth *Auth `pkl:"auth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a NatsConfig
func LoadFromPath(ctx context.Context, path string) (ret *NatsConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a NatsConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*NatsConfig, error) {
	var ret NatsConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
