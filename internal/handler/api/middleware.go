package api

import (
	"errors"
	"strings"

	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core/auth"
	"github.com/elangreza14/advance-todo/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type (
	iMiddleware interface {
		ValidateSecret(c *fiber.Ctx) error
		ValidateAccessToken(c *fiber.Ctx) error
		ValidateRefreshToken(c *fiber.Ctx) error
		OpenHandlerRateLimiter(c *fiber.Ctx) error
		AuthorizedHandlerRateLimiter(c *fiber.Ctx) error
	}

	middleware struct {
		server  Server
		conf    *config.Configuration
		service auth.AuthService
	}
)

const (
	// SecretHeaderValue is api header API for handling secret
	SecretHeaderValue string = "X-Api-Key"

	// AuthorizationHeaderValue is api header API for handling authorization bearer
	AuthorizationHeaderValue string = "Authorization"
)

var (
	// ErrMissingAPISecret is error for checking availability of the secret
	ErrMissingAPISecret = errors.New("missing api secret")

	// ErrWrongAPISecret is error for checking correctness of the secret
	ErrWrongAPISecret = errors.New("wrong api secret")

	// ErrMissingToken is error for checking availability of the token
	ErrMissingToken = errors.New("missing api token")

	// ErrWrongToken is error for checking correctness of the token
	ErrWrongToken = errors.New("wrong api token")
)

// NewMiddleware is handler for all in middleware in app
func newMiddleware(
	conf *config.Configuration,
	server Server,
	service auth.AuthService,
) iMiddleware {
	return &middleware{
		service: service,
		server:  server,
		conf:    conf,
	}
}

func (m *middleware) ValidateSecret(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	apiSecret, ok := headers[SecretHeaderValue]
	if !ok {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, ErrMissingAPISecret)
	}

	if apiSecret != m.conf.Env.XApiKey {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, ErrWrongAPISecret)
	}

	return c.Next()
}

func (m *middleware) processToken(c *fiber.Ctx) (*domain.Token, error) {
	headers := c.GetReqHeaders()
	rawToken, ok := headers[AuthorizationHeaderValue]
	if !ok {
		return nil, m.server.newErrorResponse(c, fiber.StatusUnauthorized, ErrMissingToken)
	}

	splitToken := strings.Split(rawToken, " ")
	if len(splitToken) != 2 {
		return nil, m.server.newErrorResponse(c, fiber.StatusUnauthorized, ErrWrongToken)
	}

	if splitToken[0] != "bearer" {
		return nil, m.server.newErrorResponse(c, fiber.StatusUnauthorized, ErrWrongToken)
	}

	tokenGenerator, err := m.conf.Token.Validate(splitToken[1])
	if err != nil {
		return nil, m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	m.conf.Logger.Info("", tokenGenerator.ID)

	// fill this with user id
	token, err := m.service.GetTokenByID(c.Context(), tokenGenerator.ID)
	if err != nil {
		return nil, m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	return token, nil
}

func (m *middleware) ValidateAccessToken(c *fiber.Ctx) error {
	token, err := m.processToken(c)
	if err != nil {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	if token.TokenType != domain.TokenTypeAccess {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	c.Locals(string(domain.ContextValueUserID), token.UserID.String())

	return c.Next()
}

func (m *middleware) ValidateRefreshToken(c *fiber.Ctx) error {
	token, err := m.processToken(c)
	if err != nil {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	if token.TokenType != domain.TokenTypeRefresh {
		return m.server.newErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	c.Locals(string(domain.ContextValueUserID), token.UserID.String())

	return c.Next()
}

func (m *middleware) OpenHandlerRateLimiter(c *fiber.Ctx) error {
	return c.Next()
}

func (m *middleware) AuthorizedHandlerRateLimiter(c *fiber.Ctx) error {
	return c.Next()
}
