package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-template/pkg/logging"
	"time"
)

var log = logging.MiddlewareLogger

func LoggingMiddleware(c *fiber.Ctx) error {
	//Логгирование входящих запросов
	tmStart := time.Now()
	err := c.Next()
	duration := time.Since(tmStart)

	log.Info().
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Int("status", c.Response().StatusCode()).
		Dur("duration", duration).
		Interface("queries", c.Queries()).
		Msg("Входящий запрос")

	if err != nil {
		log.Error().
			Err(err).
			Msg("Ошибка запроса")
	}
	return err
}
