package helpers

import (
	"github.com/gofiber/fiber/v2"
)

// DbQueryError - Database query error
func DbQueryError(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(fiber.Map{
		"error":   "database_query_error",
		"message": err.Error(),
	})
}

// BadRequest - Bad request
func BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(400).JSON(fiber.Map{
		"error":   "bad_request",
		"message": err.Error(),
	})
}
