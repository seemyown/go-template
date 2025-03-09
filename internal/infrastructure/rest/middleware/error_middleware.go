package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ErrorResponse struct {
	Code      int    `json:"status_code"`
	Message   string `json:"message"`
	ErrorCode string `json:"error"`
}

func (er *ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] - %s (%s)", er.Code, er.Message, er.ErrorCode)
}

func (er *ErrorResponse) Send(ctx *fiber.Ctx) error {
	return ctx.Status(er.Code).JSON(er)
}

func NewErrorResponse(err error) *ErrorResponse {
	message := exc.ErrorMessage(err)
	code := exc.ErrorStatusCode(err)
	log.Error(err, message)
	parts := strings.Split(err.Error(), ": ")
	var errCode string
	if len(parts) > 1 {
		errCode = parts[1]
	} else {
		errCode = err.Error()
	}
	response := ErrorResponse{
		Code:      code,
		Message:   message,
		ErrorCode: errCode,
	}
	return &response
}

func ErrorHandler(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err == nil {
		return nil
	}

	errStatusCode := exc.ErrorStatusCode(err)
	errMessage := exc.ErrorMessage(err)

	response := ErrorResponse{
		Code:      errStatusCode,
		Message:   errMessage,
		ErrorCode: err.Error(),
	}

	log.Error(err, "Ошибка запроса %+v", response)
	return ctx.Status(errStatusCode).JSON(response)
}
