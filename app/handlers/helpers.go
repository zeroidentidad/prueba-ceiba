package handlers

import (
	"payments/errs"

	"github.com/gofiber/fiber/v2"
)

func parseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func resJSON(c *fiber.Ctx, data interface{}, err *errs.AppError, status int) error {
	if err != nil {
		c.Status(err.Code)
		return c.JSON(&fiber.Map{
			"error": err.Message,
		})
	}

	c.Status(status)
	return c.JSON(data)
}

func setMessage(msg string) *fiber.Map {
	return &fiber.Map{
		"respuesta": msg,
	}
}
