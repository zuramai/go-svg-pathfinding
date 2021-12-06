package http

import (
	"wsc2017/app/container"
	"wsc2017/router"

	"github.com/gofiber/fiber/v2"
)

func RunServer(c *container.ServiceContainer) error {
	app := fiber.New()

	router.RegisterRoutes(app, c)

	app.Listen(":" + c.Config.AppConfig.Port)

	return nil
}
