package ext

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func (er *ErrorResponse) Send(ctx *fiber.Ctx) error {
	return ctx.Status(er.Code).JSON(er)
}
