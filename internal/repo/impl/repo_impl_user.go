package repo_impl

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/internal/entity"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"

	repo_attribute "github.com/Mind2Screen-Dev-Team/go-skeleton/internal/repo/attribute"

	"github.com/jmoiron/sqlx"
)

type userRepoImpl struct {
	db *lazy.Loader[*sqlx.DB]
}

func NewUserRepoImpl() *userRepoImpl {
	return &userRepoImpl{}
}

func (r *userRepoImpl) Loader(ctx context.Context, appDependency *registry.AppDependency, appRepository *registry.AppRepository) {
	r.db = &appDependency.MySqlDB
	appRepository.User = r
}

func (r *userRepoImpl) Find(ctx context.Context, p repo_attribute.UserFindAttribute) (*entity.User, error) {
	var u entity.User
	err := r.db.Value().GetContext(
		ctx,
		&u,
		"SELECT * FROM users WHERE (? IS NULL OR id = ?) AND (? IS NULL OR email = ?) LIMIT 1",
		p.ID,
		p.Email,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepoImpl) Count(ctx context.Context, p repo_attribute.UserFindAttribute) (int64, error) {
	var n int64
	err := r.db.Value().GetContext(
		ctx,
		&n,
		"SELECT COUNT(id) FROM users WHERE (? IS NULL OR id = ?) AND (? IS NULL OR email = ?)",
		p.ID,
		p.Email,
	)
	if err != nil {
		return 0, err
	}

	return n, nil
}
