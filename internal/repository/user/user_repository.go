package user_repo

import (
	"context"
	"go-skeleton/internal/schema"

	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository() {

}

type FindParams struct {
	ID    null.Int
	Email null.String
}

func (r *UserRepository) Find(ctx context.Context, p FindParams) (*schema.User, error) {
	var u schema.User
	err := r.db.GetContext(
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
