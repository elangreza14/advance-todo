package api

import (
	"context"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/internal/domain"
	"github.com/elangreza14/advance-todo/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type (
	iAuthAPIHandler interface {
		HandleRegister(c *fiber.Ctx) error
		HandleLogin(c *fiber.Ctx) error
		HandleGetProfile(c *fiber.Ctx) error
	}

	authAPIHandler struct {
		conf    *config.Configuration
		service auth.AuthService
		server  Server
	}
)

// newAuthHandler is handler for authentication route
func newAuthHandler(
	conf *config.Configuration,
	server Server,
	service auth.AuthService,
) iAuthAPIHandler {
	return &authAPIHandler{
		service: service,
		server:  server,
		conf:    conf,
	}
}

func (a *authAPIHandler) HandleRegister(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.RegisterUserRequest{}
	if err := a.server.bodyParser(c, req); err != nil {
		a.conf.Logger.Error("authAPIHandler.server.bodyParser", err)
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	if err := a.conf.Validator.Struct(req); err != nil {
		a.conf.Logger.Error("authAPIHandler.conf.Validator.Struct", err)
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err...)
	}

	if err := a.service.RegisterUser(contextParent, *req); err != nil {
		a.conf.Logger.Error("authAPIHandler.service.RegisterUser", err)
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusCreated, nil)
}

func (a *authAPIHandler) HandleLogin(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	req := &dto.LoginUserRequest{}
	if err := a.server.bodyParser(c, req); err != nil {
		a.conf.Logger.Error("authAPIHandler.server.bodyParser", err)
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err)
	}

	if err := a.conf.Validator.Struct(req); err != nil {
		a.conf.Logger.Error("authAPIHandler.conf.Validator.Struct", err)
		return a.server.newErrorResponse(c, fiber.StatusBadRequest, err...)
	}

	contextValue := context.WithValue(contextParent, domain.ContextValueIP, c.IP())

	res, err := a.service.LoginUser(contextValue, *req)
	if err != nil {
		a.conf.Logger.Error("authAPIHandler.service.LoginUser", err)
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusOK, res)
}

func (a *authAPIHandler) HandleGetProfile(c *fiber.Ctx) error {
	contextParent, cancel := context.WithCancel(c.Context())
	defer cancel()

	userID := c.Locals(string(domain.ContextValueUserID))

	contextValue := context.WithValue(contextParent, domain.ContextValueUserID, userID)

	res, err := a.service.GetUser(contextValue)
	if err != nil {
		a.conf.Logger.Error("authAPIHandler.service.GetUser", err)
		return a.server.newErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return a.server.newSuccessResponse(c, fiber.StatusOK, res)
}
