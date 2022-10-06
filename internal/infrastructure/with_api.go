package infrastructure

import (
	"log"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/handler/api"
)

func WithApi(env *config.Env) error {
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

	log.Fatal(app.Listen(":8080"))
	conf.Logger.Info("app is running")
	return nil
}
