package common

import "github.com/gofiber/fiber/v2"

type Response struct {
	data   map[string]interface{}
	status int
}

func NewResponse(status int, data map[string]interface{}) Response {
	return Response{
		data:   data,
		status: status,
	}
}

func (resp *Response) SendResponse(c *fiber.Ctx) error {
	c.Status(resp.status)
	return c.JSON(resp.data)
}
