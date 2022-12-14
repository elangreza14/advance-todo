package api

import (
	"github.com/gofiber/fiber/v2"
)

type (
	// Header is struct for handling all errors
	// will add something later
	Header struct {
		Errors []string `json:"errors"`
	}

	// Response is struct for wrapper all data and header
	Response struct {
		Header Header      `json:"header"`
		Data   interface{} `json:"data"`
	}
)

func (s *server) bodyParser(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return err
	}

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
