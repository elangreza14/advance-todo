package api

import (
	"context"

	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	iAuthApiHandler interface {
		HandleRegister(c *fiber.Ctx) error
		HandleLogin(c *fiber.Ctx) error
		HandleGetProfile(c *fiber.Ctx) error
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
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	if err := a.service.RegisterUser(contextParent, *req); err != nil {
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusCreated, nil)
}

func (a *authApiHandler) HandleLogin(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.LoginUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	contextValue := context.WithValue(contextParent, domain.ContextValueIP, c.IP())

	res, err := a.service.LoginUser(contextValue, *req)
	if err != nil {
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusOK, res)
}

func (a *authApiHandler) HandleGetProfile(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.LoginUserRequest{}
	if err := c.BodyParser(req); err != nil {
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	contextValue := context.WithValue(contextParent, domain.ContextValueIP, c.IP())

	res, err := a.service.LoginUser(contextValue, *req)
	if err != nil {
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusOK, res)
}
