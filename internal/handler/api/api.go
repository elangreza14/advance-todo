package api

import (
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core"

	"github.com/gofiber/fiber/v2"
)

type (
	Server interface {
		Run() error
		Shutdown() error
		newRouter()
		bodyParser(c *fiber.Ctx, data interface{}) error
		newSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error
		newErrorResponse(c *fiber.Ctx, statusCode int, dataError ...error) error
	}

	server struct {
		fbr  *fiber.App
		conf *config.Configuration
		core core.Core

		// later we add middleware here
	}

	Header struct {
		Errors []string `json:"errors"`
	}

	Response struct {
		Header Header      `json:"header"`
		Data   interface{} `json:"data"`
	}
)

func NewServer(conf *config.Configuration, core core.Core) Server {
	app := fiber.New()

	return &server{
		fbr:  app,
		conf: conf,
		core: core,
	}
}

func (s *server) Run() error {
	go func() error {
		if err := s.fbr.Listen(":8080"); err != nil {
			s.conf.Logger.Error("api.Server.Run", err)
			return err
		}
		return nil
	}()

	s.newRouter()

	return nil
}

func (s *server) Shutdown() error {
	if err := s.fbr.Shutdown(); err != nil {
		s.conf.Logger.Error("api.Server.Shutdown", err)
		return err
	}
	return nil
}

func (s *server) newRouter() {

	// auth handlers
	var handlerAuth iAuthApiHandler = NewAuthHandler(s, s.core.AuthService)
	routerAuth := s.fbr.Group("/auth")
	routerAuth.Post("/register", handlerAuth.HandleRegister)
	routerAuth.Post("/login", handlerAuth.HandleLogin)
}

func (s *server) bodyParser(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return err
	}

	// TODO implement validation

	return nil
}

func (s *server) newSuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Header: Header{
			Errors: []string{},
		},
		Data: data,
	})
}

func (s *server) newErrorResponse(c *fiber.Ctx, statusCode int, errList ...error) error {
	errs := []string{}
	for i := 0; i < len(errList); i++ {
		errs = append(errs, errList[i].Error())
	}

	return c.Status(statusCode).JSON(Response{
		Header: Header{
			Errors: errs,
		},
		Data: nil,
	})
}
