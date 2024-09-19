package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLXClient struct {
	cfg *appconfig.AppConfig
}

func NewMySQLX(cfg *appconfig.AppConfig) *MySQLXClient {
	return &MySQLXClient{cfg}
}

func (s *MySQLXClient) Create(_ context.Context) (db *sqlx.DB, err error) {
	if !s.cfg.Mysql.Enabled {
		return nil, errors.New("mysql client database is disabled")
	}

	db, err = sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			s.cfg.Mysql.Auth.Username,
			s.cfg.Mysql.Auth.Password,
			s.cfg.Mysql.Host,
			s.cfg.Mysql.Port,
			s.cfg.Mysql.Db,
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

func (s *MySQLXClient) Loader(ctx context.Context, app *registry.AppDependency) {
	app.MySqlDB = lazy.New(func() (db *sqlx.DB, err error) {
		return s.Create(ctx)
	})
}
