// Code generated from Pkl module `MinioConfig`. DO NOT EDIT.
package minioconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type MinioConfig struct {
	// Minio feature is enabled?
	Enabled bool `pkl:"enabled"`

	// The endpoint that minio cluster
	Endpoint string `pkl:"endpoint"`

	// Minio options secure is enabled?
	UseSSL bool `pkl:"useSSL"`

	// Credential settings for Minio
	Credential *Credential `pkl:"credential"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a MinioConfig
func LoadFromPath(ctx context.Context, path string) (ret *MinioConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a MinioConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*MinioConfig, error) {
	var ret MinioConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
