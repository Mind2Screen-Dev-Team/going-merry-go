package registry

import (
	service_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/api"
)

// # SERVICE

type ServiceRegistry struct {
	// register your service on here
	User service_api.UserServiceAPI
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{}
}
