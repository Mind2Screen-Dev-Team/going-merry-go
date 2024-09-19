package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	cfg *appconfig.AppConfig
}

func NewRedisClient(cfg *appconfig.AppConfig) *RedisClient {
	return &RedisClient{cfg}
}

func (r *RedisClient) Create(_ context.Context) (*redis.Client, error) {
	if !r.cfg.Redis.Enabled {
		return nil, errors.New("redis client is disabled")
	}

	op := redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.cfg.Redis.Host, r.cfg.Redis.Port),
		DB:   r.cfg.Redis.Db,
	}

	if r.cfg.Redis.Auth.Enabled {
		op.Username = r.cfg.Redis.Auth.Username
		op.Password = r.cfg.Redis.Auth.Password
	}

	return redis.NewClient(&op), nil
}

func (r *RedisClient) Loader(ctx context.Context, app *bootstrap.AppDependency) {
	app.Redis = lazy.New(func() (*redis.Client, error) {
		return r.Create(ctx)
	})
}
