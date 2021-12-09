package errors

import (
	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(c *fiber.Ctx, err ErrorResponse) error {
	return c.JSON(err)
}
