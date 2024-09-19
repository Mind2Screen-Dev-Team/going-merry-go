// Code generated from Pkl module `MySQLConfig`. DO NOT EDIT.
package mysqlconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type MySQLConfig struct {
	// MySql Feature is enabled
	Enabled bool `pkl:"enabled"`

	// The hostname that MySQL listens on
	Host string `pkl:"host"`

	// The port that MySQL listens on
	Port int `pkl:"port"`

	// The port that MySQL listens on
	Db string `pkl:"db"`

	// Authorization settings for MySQL
	Auth *Auth `pkl:"auth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a MySQLConfig
func LoadFromPath(ctx context.Context, path string) (ret *MySQLConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a MySQLConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*MySQLConfig, error) {
	var ret MySQLConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
