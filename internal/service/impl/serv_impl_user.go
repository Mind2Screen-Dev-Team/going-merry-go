package service_impl

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/entity"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"
	"github.com/redis/go-redis/v9"

	repo_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/api"
	repo_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/attribute"
	service_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/attribute"
)

type userServiceImpl struct {
	redis *lazy.Loader[*redis.Client]
	user  repo_api.UserRepoAPI
}

func NewUserServiceImpl() *userServiceImpl {
	return &userServiceImpl{}
}

func (r *userServiceImpl) Loader(ctx context.Context, appDependency *bootstrap.AppDependency, appRepository *bootstrap.AppRepository, appService *bootstrap.AppService) {
	r.redis = &appDependency.Redis
	appService.User = r
}

func (r *userServiceImpl) Find(ctx context.Context, p service_attribute.UserFindAttribute) (*entity.User, error) {
	return r.user.Find(ctx, repo_attribute.UserFindAttribute{
		ID:    p.ID,
		Email: p.Email,
	})
}
