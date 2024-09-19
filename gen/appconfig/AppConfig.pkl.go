// Code generated from Pkl module `AppConfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/mysqlconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/natsconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/redisconfig"
	"github.com/apple/pkl-go/pkl"
)

type AppConfig struct {
	// App Name
	AppName string `pkl:"appName"`

	// App Domain
	AppDomain string `pkl:"appDomain"`

	// App Host
	AppHost string `pkl:"appHost"`

	// App JWT Secret
	AppJwtSecret string `pkl:"appJwtSecret"`

	// App Log Path
	AppLogPath string `pkl:"appLogPath"`

	// App Http Configuration
	AppHttp *AppHttp `pkl:"appHttp"`

	// MySql Configuration
	Mysql *mysqlconfig.MySQLConfig `pkl:"mysql"`

	// Redis Configuration
	Redis *redisconfig.RedisConfig `pkl:"redis"`

	// Nats Configuration
	Nats *natsconfig.NatsConfig `pkl:"nats"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *AppConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AppConfig, error) {
	var ret AppConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
