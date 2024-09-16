package router

import (
	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
)

type AuthRouter struct {
	cfg        *appconfig.AppConfig
	dependency *bootstrap.AppDependency
	repo       *bootstrap.AppRepository
	service    *bootstrap.AppService
}

func (r *AuthRouter) Loader(cfg *appconfig.AppConfig, dep *bootstrap.AppDependency, repo *bootstrap.AppRepository, service *bootstrap.AppService) {
	r.cfg = cfg
	r.dependency = dep
	r.repo = repo
	r.service = service
}
