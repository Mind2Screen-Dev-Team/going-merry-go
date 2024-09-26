package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/app/registry"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/xlazy"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mySqlXClient struct{}

func NewMySqlX() *mySqlXClient {
	return &mySqlXClient{}
}

func (s *mySqlXClient) Create(_ context.Context, cfg *appconfig.AppConfig) (db *sqlx.DB, err error) {
	if !cfg.Mysql.Enabled {
		return nil, errors.New("mysql client database is disabled")
	}

	db, err = sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Mysql.Auth.Username,
			cfg.Mysql.Auth.Password,
			cfg.Mysql.Host,
			cfg.Mysql.Port,
			cfg.Mysql.Db,
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

func (s *mySqlXClient) Loader(ctx context.Context, cfg *appconfig.AppConfig, app *registry.AppDependency) {
	app.MySqlDB = xlazy.New(func() (db *sqlx.DB, err error) {
		return s.Create(ctx, cfg)
	})
}
