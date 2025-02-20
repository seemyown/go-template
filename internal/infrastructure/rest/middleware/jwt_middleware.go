package middleware

import (
	"errors"
	"fmt"
	"go-fiber-template/pkg/config"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warn().Msg("Не авторизованный запрос")
			return c.Next()
		}
		tokenStr, err := extractToken(authHeader)
		if err != nil {
			log.Warn().Msg("Не верно передан токен")
			return c.Next()
		}
		log.Info().Msgf("Входящий запрос с токеном : %s", tokenStr)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(config.SecretKey), nil
		})

		if err != nil || !token.Valid {
			log.Error().Err(err).Msg("Некорректный токен")
			return err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Error().Err(err).Msg("Некорректный токен")
			return err
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			log.Error().Err(err).Msg("Неверный токен")
			return err
		}
		if int64(exp) < time.Now().Unix() {
			log.Error().Err(err).Msg("Токен просрочен")
			return err
		}

		iss, ok := claims["iss"].(string)
		if !ok {
			log.Error().Err(err).Msg("Неверный токен")
			return err
		}
		if iss != config.TokenIssuer {
			log.Error().Err(err).Msg("Неверный источник токена")
			return err
		}
		// Прочая логика обработки
		return c.Next()
	}
}

func extractToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}
