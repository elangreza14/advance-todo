package infrastructure

import (
	"context"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/handler/api"
)

func Run(env *config.Env) error {
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
		config.WithRedis(),
	)

	app := api.New(conf)

	go func() error {
		if err = app.Listen(":8080"); err != nil {
			conf.Logger.Error("main.err", err)
			return err
		}
		return nil
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
