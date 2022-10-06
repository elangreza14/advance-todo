package infrastructure

import (
	"log"
	"os"
	"time"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/internal/handler/api"
	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"
	"github.com/gofiber/fiber/v2"
)

func WithApi(env *config.Env) {
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
		os.Exit(1)
	}

	postgresRepository := postgresRepo.NewPostgresRepo(conf, postgresRepo.WithUser())

	authService := auth.NewAuthService(conf, postgresRepository.User)

	app := fiber.New()

	api.NewAuthApiHandler(conf, authService, app.Group("/auth"))

	log.Fatal(app.Listen(":8080"))
	conf.Logger.Info("app is running")
}
