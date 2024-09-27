// Code generated from Pkl module `NatsConfig`. DO NOT EDIT.
package natsconfig

type Auth struct {
	// Nats authorization is nabled
	Enabled bool `pkl:"enabled"`

	// Nats authorization username
	Username string `pkl:"username"`

	// Nats authorization password
	Password string `pkl:"password"`
}
