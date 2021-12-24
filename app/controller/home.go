package controller

import "github.com/gofiber/fiber/v2"

func (controller *Controller) Home(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Welcome to Go Pathfinding API",
	})
}
