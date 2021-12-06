package controller

import "wsc2017/app/container"

type Controller struct {
	Container *container.ServiceContainer
}

func NewController(c *container.ServiceContainer) *Controller {
	controller := new(Controller)
	controller.Container = c

	return controller
}
