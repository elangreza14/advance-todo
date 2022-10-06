package api

import (
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core"

	postgresRepo "github.com/elangreza14/advance-todo/adapter/postgres"

	"github.com/gofiber/fiber/v2"
)

func New(conf *config.Configuration) *fiber.App {
	postgresRepository := postgresRepo.NewPostgresRepo(conf, postgresRepo.WithUser())
	newService := core.New(conf, postgresRepository)

	app := fiber.New()
	newAuthApiHandler(conf, newService.AuthService, app.Group("/auth"))
	return app
}
