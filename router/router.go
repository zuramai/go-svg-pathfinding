package router

import (
	"wsc2017/app/container"
	"wsc2017/app/controller"
	"wsc2017/app/middleware"
	"wsc2017/internal/common"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, sc *container.ServiceContainer) {
	c := controller.NewController(sc)
	m := middleware.NewMiddleware(sc)

	app.Get("/", c.Home)

	api := app.Group("/v1")
	api.Post("/auth/login", c.Login)

	app.Use(m.AuthMiddleware)
	api.Get("/auth/logout", c.Logout)

	api.Get("/place", c.GetPlace)
	api.Post("/place", c.StorePlace)
	api.Get("/place/:id", c.ShowPlace)
	api.Delete("/place/:id", c.DeletePlace)
	api.Put("/place/:id", c.UpdatePlace)

	api.Post("/schedule", c.StoreSchedule)

	app.Static("/storage", "../storage")

	app.Use(func(c *fiber.Ctx) error {
		response := common.NewResponse(404, fiber.Map{
			"message": "Not Found",
		})
		return response.SendResponse(c)
	})
}
