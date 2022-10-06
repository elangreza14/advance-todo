package main

import (
	"log"
	"os"
	"time"

	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/api"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatal(err)
	}

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
