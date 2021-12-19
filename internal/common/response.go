package common

import "github.com/gofiber/fiber/v2"

type Response struct {
	data map[string]interface{}
}

func NewResponse(data map[string]interface{}) Response {
	return Response{
		data: data,
	}
}

func (resp *Response) SendResponse(ctx *fiber.Ctx) error {
	return ctx.JSON(resp.data)
}
