package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go-fiber-template/pkg/logging"
	"strings"
	"time"
)

var _log = logging.New(logging.Config{
	FileName: "jwt_middleware.log",
	Name:     "jwt_middleware",
})

func JWTMiddleware(
	secret string,
	checks []func(claims jwt.MapClaims) error,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			_log.Warn("Не авторизованный запрос")
			return c.Next()
		}
		tokenStr, err := extractToken(authHeader)
		if err != nil {
			_log.Warn("Не верно передан токен")
			return c.Next()
		}
		_log.Info("Входящий запрос с токеном : %s", tokenStr)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			_log.Error(err, "Некорректный токен")
			return err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			_log.Error(err, "Некорректный токен")
			return err
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			_log.Error(err, "Неверный токен")
			return err
		}
		if int64(exp) < time.Now().Unix() {
			_log.Error(err, "Токен просрочен")
			return err
		}

		for _, check := range checks {
			if err = check(claims); err != nil {
				_log.Error(err, "Ошибка валидации токена")
				return err
			}
		}

		payload, ok := claims["payload"].(map[string]interface{})
		if !ok {
			_log.Error(err, "Не верное содержимое токена")
			return err
		}
		for k, v := range payload {
			c.Locals(k, v)
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
