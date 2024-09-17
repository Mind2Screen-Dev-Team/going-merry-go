package repo_impl

import (
	"context"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/bootstrap"
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

func (r *userRepoImpl) Loader(ctx context.Context, appDependency *bootstrap.AppDependency, appRepository *bootstrap.AppRepository) {
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
