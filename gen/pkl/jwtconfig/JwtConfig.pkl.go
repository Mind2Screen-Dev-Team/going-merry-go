// Code generated from Pkl module `JwtConfig`. DO NOT EDIT.
package jwtconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type JwtConfig struct {
	// JWT Secret Key
	Secret string `pkl:"secret"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a JwtConfig
func LoadFromPath(ctx context.Context, path string) (ret *JwtConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a JwtConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*JwtConfig, error) {
	var ret JwtConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
