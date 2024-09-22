package registry

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

// # DEPENDENCY

type AppDependency struct {
	// register your dependency on here
	Logger            *zerolog.Logger
	MySqlDB           lazy.Loader[*sqlx.DB]
	Redis             lazy.Loader[*redis.Client]
	NatsConn          lazy.Loader[*nats.Conn]
	NatsJetStreamConn lazy.Loader[jetstream.JetStream]
}
