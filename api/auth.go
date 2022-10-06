package api

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthApiHandler interface {
		Register(c *fiber.Ctx) error
	}

	authApiHandler struct {
		conf    *config.Configuration
		service auth.AuthService
	}
)

func newHandler(conf *config.Configuration, service auth.AuthService) AuthApiHandler {
	return &authApiHandler{
		conf:    conf,
		service: service,
	}
}

func NewAuthApiHandler(conf *config.Configuration, service auth.AuthService, router fiber.Router) {
	a := newHandler(conf, service)

	router.Post("/register", a.Register)
}

func (a *authApiHandler) Register(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(context.Background())
	defer cancel()

	req := &dto.RegisterUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
	}

	if err := a.service.RegisterUser(contextParent, *req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).JSON(
		&fiber.Map{
			"status":  "success",
			"message": "created",
		})
}
