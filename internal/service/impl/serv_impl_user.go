package service_impl

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/entity"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"
	"github.com/redis/go-redis/v9"

	repo_api "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/api"
	repo_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/attribute"
	service_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/service/attribute"
)

type userServiceImpl struct {
	redis *xlazy.Loader[*redis.Client]
	user  repo_api.UserRepoAPI
}

func NewUserServiceImpl() *userServiceImpl {
	return &userServiceImpl{}
}

func (r *userServiceImpl) Loader(ctx context.Context, reg *registry.AppRegistry) error {
	r.redis = &reg.Dependency.Redis
	r.user = reg.Repository.User
	reg.Service.User = r
	return nil
}

func (r *userServiceImpl) Find(ctx context.Context, p service_attribute.UserFindAttribute) (*entity.User, error) {
	return r.user.Find(ctx, repo_attribute.UserFindAttribute{
		ID:    p.ID,
		Email: p.Email,
	})
}
