package fiber_api

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

type (
	IAuthApiHandler interface {
		Register(c *fiber.Ctx) error
	}

	authApiHandler struct {
		conf    *config.Configuration
		service auth.AuthService
	}
)

func NewAuthFiber(conf *config.Configuration, service auth.AuthService) IAuthApiHandler {
	return &authApiHandler{
		conf:    conf,
		service: service,
	}
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
