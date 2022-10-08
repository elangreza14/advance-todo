package infrastructure

import (
	"context"
	"log"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/handler/api"
)

func RunHttp(env *config.Env) error {
	conf, err := config.NewConfig(
		env,
		config.WithPostgres(config.DbSqlOption{
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
	)

	if err != nil {
		conf.Logger.Error("main.err", err)
		return err
	}

	app := api.New(conf)

	go func() {
		if err = app.Listen(":8080"); err != nil {
			conf.Logger.Error("main.err", err)
			log.Fatal(err)
		}
	}()

	conf.Logger.Info("app is running")

	GracefulShutdown(JobFunction{
		"fiber": func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				return err
			}
			return nil
		},
		"postgres": func(ctx context.Context) error {
			if err := conf.DbSql.Close(); err != nil {
				return err
			}
			return nil
		},
	})

	return nil
}
