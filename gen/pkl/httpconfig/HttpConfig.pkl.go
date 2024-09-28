// Code generated from Pkl module `HttpConfig`. DO NOT EDIT.
package httpconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type HttpConfig struct {
	// Service HTTP Name
	ServiceName string `pkl:"serviceName"`

	// Service HTTP Domain
	Domain string `pkl:"domain"`

	// Service HTTP Host
	Host string `pkl:"host"`

	// Service HTTP Port
	Port int `pkl:"port"`

	// Service HTTP Idle Timeout
	IdleTimeout int `pkl:"idleTimeout"`

	// Service HTTP Read Header Timeout
	ReadHeaderTimeout int `pkl:"readHeaderTimeout"`

	// Service HTTP Read Timeout
	ReadTimeout int `pkl:"readTimeout"`

	// Service HTTP Write Timeout
	WriteTimeout int `pkl:"writeTimeout"`

	// Service HTTP Handler Timeout
	HandlerTimeout int `pkl:"handlerTimeout"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a HttpConfig
func LoadFromPath(ctx context.Context, path string) (ret *HttpConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a HttpConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*HttpConfig, error) {
	var ret HttpConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
