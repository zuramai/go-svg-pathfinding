package errors

import (
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(c *fiber.Ctx, err error) error {
	serr, ok := err.(*ErrorResponse)
	if !ok {
		return c.JSON(fiber.Map{"error": err.Error()})
	}
	logger.SugarLog.Debug("statusCode %d", serr.statusCode)
	return c.Status(serr.statusCode).JSON(serr)
}
