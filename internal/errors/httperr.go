package errors

import (
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(c *fiber.Ctx, err error) error {
	serr, ok := err.(*ErrorResponse)
	logger.SugarLog.Debug("statusCode %d", serr.statusCode)
	if !ok {
		return c.JSON(serr)
	}
	return c.Status(serr.statusCode).JSON(serr)
}
