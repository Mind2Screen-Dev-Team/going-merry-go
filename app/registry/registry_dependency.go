package registry

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/redis/go-redis/v9"
)

// # DEPENDENCY

type AppDependency struct {
	// register your dependency on here
	ZeroLogDefaultFields map[string]any
	LumberjackLogger     *lumberjack.Logger
	ZeroLogger           zerolog.Logger
	MySqlDB              xlazy.Loader[*sqlx.DB]
	Redis                xlazy.Loader[*redis.Client]
	NatsConn             xlazy.Loader[*nats.Conn]
	NatsJetStreamConn    xlazy.Loader[jetstream.JetStream]
}
