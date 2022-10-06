package api

import (
	fiberApi "github.com/elangreza14/advance-todo/adapter/fiber_api"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthApiHandler interface {
		Register(c *fiber.Ctx) error
	}
)

func NewAuthApiHandler(conf *config.Configuration, service auth.AuthService, router fiber.Router) {
	var h AuthApiHandler = fiberApi.NewAuthFiber(conf, service)
	router.Post("/register", h.Register)
}
