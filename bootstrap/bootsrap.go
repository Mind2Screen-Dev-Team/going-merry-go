package bootstrap

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

type LoaderFn interface {
	Loader(ctx context.Context, dep *Depedency)
}

type Depedency struct {
	MySqlDB           lazy.Loader[*sqlx.DB]
	Redis             lazy.Loader[*redis.Client]
	NatsConn          lazy.Loader[*nats.Conn]
	NatsJetStreamConn lazy.Loader[*nats.JetStreamContext]
}

func Load(ctx context.Context, loaders ...LoaderFn) *Depedency {
	var dep Depedency

	if loaders == nil {
		return &dep
	}

	for _, l := range loaders {
		l.Loader(ctx, &dep)
	}

	return &dep
}
