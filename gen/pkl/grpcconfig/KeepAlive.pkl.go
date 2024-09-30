// Code generated from Pkl module `GrpcConfig`. DO NOT EDIT.
package grpcconfig

// Service GRPC Keep Alive configuration
type KeepAlive struct {
	Enabled bool `pkl:"enabled"`

	Policy *KeepAliveEnforcementPolicy `pkl:"policy"`

	Parameter *KeepAliveServerParameter `pkl:"parameter"`
}
