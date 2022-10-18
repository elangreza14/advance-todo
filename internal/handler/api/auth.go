package api

import (
	"context"

	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	iAuthApiHandler interface {
		HandleRegister(c *fiber.Ctx) error
		HandleLogin(c *fiber.Ctx) error
	}

	authApiHandler struct {
		service auth.AuthService
		server  Server
	}
)

func NewAuthHandler(server Server, service auth.AuthService) iAuthApiHandler {
	return &authApiHandler{
		service: service,
		server:  server,
	}
}

func (a *authApiHandler) HandleRegister(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.RegisterUserRequest{}
	if err := a.server.bodyParser(c, req); err != nil {
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

func (a *authApiHandler) HandleLogin(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.LoginUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
	}

	res, err := a.service.LoginUser(contextParent, *req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
