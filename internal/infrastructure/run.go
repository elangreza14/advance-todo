package infrastructure

import (
	"time"

	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core"
	"github.com/elangreza14/advance-todo/internal/handler/api"
)

func Run(env *config.Env) error {
	conf, err := config.NewConfig(
		env,
		config.WithDBSql(config.DbSqlOption{
			MaxLifeTime:  time.Duration(5 * time.Minute),
			MaxIdleConns: 25,
			MaxOpenConns: 25,
		}),
		config.WithLogger(config.LoggerOption{
			EncodingType:     config.EncodingTypeConsole,
			NameService:      "advance-todo",
			EnableStackTrace: false,
			IsDebug:          false,
		}),
		config.WithCache(),
		config.WithToken(),
	)
	if err != nil {
		return err
	}

	postgresRepository := postgresRepo.NewPostgresRepo(conf, postgresRepo.WithUser(), postgresRepo.WithToken())
	coreApp := core.New(conf, postgresRepository)
	app := api.NewServer(conf, coreApp)

	app.Run()

	conf.Logger.Info("app is running")

	GracefulShutdown(JobFunction{
		"fiber": func() error {
			if err := app.Shutdown(); err != nil {
				return err
			}
			return nil
		},
		"postgres": func() error {
			if err := conf.DbSql.Close(); err != nil {
				return err
			}
			return nil
		},
	})

	return nil
}
