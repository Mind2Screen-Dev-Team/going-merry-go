package repo_api

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/entity"
	repo_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/attribute"
)

type UserRepoAPI interface {
	Find(ctx context.Context, p repo_attribute.UserFindAttribute) (*entity.User, error)
}
