package config

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"
)

type NatsClient struct {
	appConfig *appconfig.AppConfig
}

func NewNatsClient(appConfig *appconfig.AppConfig) *NatsClient {
	return &NatsClient{appConfig}
}

func (n *NatsClient) Create(_ context.Context) (*nats.Conn, error) {
	var options []nats.Option
	if n.appConfig.Nats.Enabled {
		options = append(options, nats.UserInfo(
			n.appConfig.Nats.Auth.Username,
			n.appConfig.Nats.Auth.Password,
		))
	}

	return nats.Connect(
		fmt.Sprintf(
			"nats://%s:%d",
			n.appConfig.Nats.Host,
			n.appConfig.Nats.Port,
		),
		options...,
	)
}

func (n *NatsClient) Loader(ctx context.Context, app *bootstrap.AppDependency) {
	app.NatsConn = lazy.New(func() (*nats.Conn, error) {
		return n.Create(ctx)
	})

	app.NatsJetStreamConn = lazy.New(func() (jetstream.JetStream, error) {
		return jetstream.New(app.NatsConn.Value())
	})
}
