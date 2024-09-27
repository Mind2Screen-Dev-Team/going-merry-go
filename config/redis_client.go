package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/pkl/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"
	"github.com/redis/go-redis/v9"
)

type redisClient struct{}

func NewRedisClient() *redisClient {
	return &redisClient{}
}

func (r *redisClient) Create(_ context.Context, cfg *appconfig.AppConfig) (*redis.Client, error) {
	if !cfg.Redis.Enabled {
		return nil, errors.New("redis client is disabled")
	}

	op := redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.Db,
	}

	if cfg.Redis.Auth.Enabled {
		op.Username = cfg.Redis.Auth.Username
		op.Password = cfg.Redis.Auth.Password
	}

	return redis.NewClient(&op), nil
}

func (r *redisClient) Loader(ctx context.Context, cfg *appconfig.AppConfig, app *registry.AppDependency) {
	app.Redis = xlazy.New(func() (*redis.Client, error) {
		return r.Create(ctx, cfg)
	})
}
