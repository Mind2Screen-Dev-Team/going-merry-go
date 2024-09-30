// Code generated from Pkl module `GrpcConfig`. DO NOT EDIT.
package grpcconfig

// Service GRPC Keep Alive Server Parameters
type KeepAliveServerParameter struct {
	// If a client is idle for X seconds, send a GOAWAY
	MaxConnectionIdle int `pkl:"maxConnectionIdle"`

	// If any connection is alive for more than X seconds, send a GOAWAY
	MaxConnectionAge int `pkl:"maxConnectionAge"`

	// Allow X seconds for pending RPCs to complete before forcibly closing connections
	MaxConnectionAgeGrace int `pkl:"maxConnectionAgeGrace"`

	// Ping the client if it is idle for X seconds to ensure the connection is still active
	Time int `pkl:"time"`

	// Wait X second for the ping ack before assuming the connection is dead
	Timeout int `pkl:"timeout"`
}
