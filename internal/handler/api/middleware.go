package api

import (
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/gofiber/fiber/v2"
)

type (
	iMiddleware interface {
		ValidateSecret(c *fiber.Ctx) error
		ValidateBearerToken(c *fiber.Ctx) error
		OpenHandlerRateLimiter(c *fiber.Ctx) error
		AuthorizedHandlerRateLimiter(c *fiber.Ctx) error
	}

	middleware struct {
		server Server
	}
)

// NewMiddleware is handler for all in middleware in app
func NewMiddleware(server Server, service auth.AuthService) iMiddleware {
	return &middleware{
		server: server,
	}
}

func (m *middleware) ValidateSecret(c *fiber.Ctx) error {
	return c.Next()
}

func (m *middleware) ValidateBearerToken(c *fiber.Ctx) error {
	return c.Next()
}

func (m *middleware) OpenHandlerRateLimiter(c *fiber.Ctx) error {
	return c.Next()
}

func (m *middleware) AuthorizedHandlerRateLimiter(c *fiber.Ctx) error {
	return c.Next()
}
