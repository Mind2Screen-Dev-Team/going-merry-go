package config

import (
	"context"
	"fmt"
	"go-skeleton/bootstrap"
	"go-skeleton/pkg/lazy"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLX struct {
	dsn DSN
}

func NewMySQLX(dsn DSN) *MySQLX {
	return &MySQLX{dsn}
}

func (s *MySQLX) Create(_ context.Context) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			s.dsn.User, s.dsn.Pass, s.dsn.Host, s.dsn.Port, s.dsn.DB,
		),
	)
	if err != nil {
		defer db.Close()
		return
	}

	err = db.Ping()
	if err != nil {
		defer db.Close()
		return
	}

	return db, err
}

func (s *MySQLX) Loader(ctx context.Context, app *bootstrap.Depedency) {
	app.MySqlDB = lazy.New(func() (db *sqlx.DB, err error) {
		return s.Create(ctx)
	})
}
