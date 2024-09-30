// Code generated from Pkl module `GrpcConfig`. DO NOT EDIT.
package grpcconfig

// Service GRPC Keep Alive Enforcement Policy
type KeepAliveEnforcementPolicy struct {
	// If a client pings more than once every X seconds, terminate the connection
	MinTime int `pkl:"minTime"`

	// Allow pings even when there are no active streams
	PermitWithoutStream bool `pkl:"permitWithoutStream"`
}
