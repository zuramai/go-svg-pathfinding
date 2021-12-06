package router

import (
	"wsc2017/app/container"
	"wsc2017/app/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, sc *container.ServiceContainer) {
	c := controller.NewController(sc)

	api := app.Group("/api/v1")
	api.Post("/auth/login", c.Login)
	api.Get("/auth/logout", c.Logout)

	api.Get("/place", c.GetPlace)
	api.Post("/place", c.StorePlace)
	api.Get("/place/:id", c.ShowPlace)
	api.Delete("/place/:id", c.DeletePlace)
	api.Put("/place/:id", c.UpdatePlace)

	api.Post("/schedule", c.StoreSchedule)
}
