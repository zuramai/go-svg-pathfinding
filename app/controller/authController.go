package controller

import "github.com/gofiber/fiber/v2"

func (controller *Controller) Login(c *fiber.Ctx) error {
	c.Send([]byte("bisa anjir"))
	return nil
}

func (controller *Controller) Logout(c *fiber.Ctx) error {
	return nil
}
