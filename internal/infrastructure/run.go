package infrastructure

import (
	"log"
	"time"

	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core"
	"github.com/elangreza14/advance-todo/internal/handler/api"
)

// Run is function to start all api
func Run(env *config.Env) error {
	conf, err := config.NewConfig(
		env,
		config.WithDBSql(config.DbSQLOption{
			MaxLifeTime:  5 * time.Minute,
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
		config.WithValidator(),
	)
	if err != nil {
		return err
	}

	postgresRepository := postgresRepo.NewPostgresRepo(conf, postgresRepo.WithUser(), postgresRepo.WithToken())
	coreApp := core.New(conf, postgresRepository)
	app := api.NewServer(conf, coreApp)

	go func() {
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	conf.Logger.Info("app is running")

	GracefulShutdown(JobFunction{
		"fiber": func() error {
			if err := app.Shutdown(); err != nil {
				return err
			}
			return nil
		},
		"postgres": func() error {
			if err := conf.DbSQL.Close(); err != nil {
				return err
			}
			return nil
		},
	})

	return nil
}
