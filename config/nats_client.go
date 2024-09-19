package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"
)

type NatsClient struct {
	cfg *appconfig.AppConfig
}

func NewNatsClient(cfg *appconfig.AppConfig) *NatsClient {
	return &NatsClient{cfg}
}

func (n *NatsClient) Create(_ context.Context) (*nats.Conn, error) {
	if !n.cfg.Nats.Enabled {
		return nil, errors.New("nats client message broker is disabled")
	}

	var options []nats.Option
	if n.cfg.Nats.Auth.Enabled {
		options = append(options, nats.UserInfo(
			n.cfg.Nats.Auth.Username,
			n.cfg.Nats.Auth.Password,
		))
	}

	return nats.Connect(
		fmt.Sprintf(
			"nats://%s:%d",
			n.cfg.Nats.Host,
			n.cfg.Nats.Port,
		),
		options...,
	)
}

func (n *NatsClient) Loader(ctx context.Context, app *registry.AppDependency) {
	app.NatsConn = lazy.New(func() (*nats.Conn, error) {
		return n.Create(ctx)
	})

	app.NatsJetStreamConn = lazy.New(func() (jetstream.JetStream, error) {
		return jetstream.New(app.NatsConn.Value())
	})
}
