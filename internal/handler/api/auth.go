package api

import (
	fiberApi "github.com/elangreza14/advance-todo/adapter/fiber_api"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

type (
	authApiHandler interface {
		Register(c *fiber.Ctx) error
	}
)

func newAuthApiHandler(conf *config.Configuration, service auth.AuthService, router fiber.Router) {
	var h authApiHandler = fiberApi.NewAuthFiber(conf, service)

	router.Post("/register", h.Register)
}
