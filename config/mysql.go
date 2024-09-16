package config

import (
	"context"
	"fmt"

	"github.com/Mind2Screen-Dev-Team/go-skeleton/bootstrap"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/gen/appconfig"
	"github.com/Mind2Screen-Dev-Team/go-skeleton/pkg/lazy"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLX struct {
	appConfig *appconfig.AppConfig
}

func NewMySQLX(appConfig *appconfig.AppConfig) *MySQLX {
	return &MySQLX{appConfig}
}

func (s *MySQLX) Create(_ context.Context) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			s.appConfig.Mysql.Auth.Username,
			s.appConfig.Mysql.Auth.Password,
			s.appConfig.Mysql.Host,
			s.appConfig.Mysql.Port,
			s.appConfig.Mysql.Db,
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

func (s *MySQLX) Loader(ctx context.Context, app *bootstrap.AppDependency) {
	app.MySqlDB = lazy.New(func() (db *sqlx.DB, err error) {
		return s.Create(ctx)
	})
}
