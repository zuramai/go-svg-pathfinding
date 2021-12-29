package middleware

import "wsc2017/app/container"

type Middleware struct {
	container *container.ServiceContainer
}

func NewMiddleware(container *container.ServiceContainer) *Middleware {
	return &Middleware{
		container: container,
	}
}
