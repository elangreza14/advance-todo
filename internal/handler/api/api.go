package api

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/elangreza14/advance-todo/config"
	"github.com/elangreza14/advance-todo/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type (

	// Server is interface that holding all server behavior
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
)

// NewServer is a server wrapper
// you can run and shutdown
func NewServer(conf *config.Configuration, core core.Core) Server {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, https://gofiber.net", // add later
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// TODO do logger based on env condition
	app.Use(logger.New(logger.Config{
		Format:     "${time} | [${ip}]:${port} ${status} - ${method} ${path} \n",
		TimeFormat: time.RFC1123,
		TimeZone:   "Asia/Jakarta",
	}))

	return &server{
		fbr:  app,
		conf: conf,
		core: core,
	}
}

func (s *server) Run() error {
	s.newRouter()

	if err := s.fbr.Listen(":8080"); err != nil {
		s.conf.Logger.Error("server.fbr.Listen", err)
		return err
	}

	return nil
}

func (s *server) Shutdown() error {
	if err := s.fbr.Shutdown(); err != nil {
		s.conf.Logger.Error("server.fbr.Shutdown", err)
		return err
	}
	return nil
}
